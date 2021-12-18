import codecs
import socket
import os
from pdfminer.pdfinterp import PDFResourceManager, PDFPageInterpreter
from pdfminer.converter import TextConverter
from pdfminer.layout import LAParams
from pdfminer.pdfpage import PDFPage
from io import StringIO
from PIL import Image
from io import BytesIO
import string
import hw1_utils
# from pdfminer import high_level
import urllib
import webbrowser

# Define socket host and port
SERVER_HOST = '127.0.0.1'
SERVER_PORT = 8888

def filter_stopword(txt_file):
    txt_file = txt_file.replace('\n', " ")
    txt_array = txt_file.split(" ")
    # print(txt_array)
    matching_words = ""
    a_file = open("stopwords.txt", "r")

    list_of_lists = []
    for line in a_file:
        stripped_line = line.strip()
        line_list = stripped_line.split()
        list_of_lists.append(line_list)

    a_file.close()

    for i in range(len(list_of_lists)):
            if list_of_lists[i][0] in txt_array:
                txt_array = list(filter(lambda a: a != (list_of_lists[i][0]), txt_array))
    str1 = ' '.join(txt_array)
    return str1


def convert_pdf_to_txt(path):
    rsrcmgr = PDFResourceManager()
    retstr = StringIO()
    codec = 'utf-8'
    laparams = LAParams()
    device = TextConverter(rsrcmgr, retstr, laparams=laparams)
    fp = open(path, 'rb')
    interpreter = PDFPageInterpreter(rsrcmgr, device)
    password = ""
    maxpages = 0
    caching = True
    pagenos = set()

    for page in PDFPage.get_pages(fp, pagenos, maxpages=maxpages, password=password, caching=caching,
                                  check_extractable=True):
        interpreter.process_page(page)

    text = retstr.getvalue()

    fp.close()
    device.close()
    retstr.close()
    return text.lower()


def file_runner():
    pdf_texts = []
    files_list = []
    files_paths_list = []
    files_paths_dict = {}
    for root, dirs, files in os.walk("pdfs"):
        path = root.split(os.sep)
        for file in files:
            files_list.append(file)
            s = "/"
            s = s.join(path)
            # adding filePath
            rootPathSplit = root.split('\\')
            rootPath = ''
            for rs in rootPathSplit:
                rootPath += rs
                rootPath += '/'
            filePath = rootPath + file
            # print(filePath)
            files_paths_list.append(filePath)
            files_paths_dict[filePath] = file

            pdf_texts.append(convert_pdf_to_txt(s + "/" + file))
    # print(pdf_texts)
    # print(files_list)
    return [files_paths_list, files_paths_dict]


def get_home_page():
    send_data = "HTTP/1.1 200 OK \r\n"
    send_data += "Content-Type: text/html; charset=utf-8\r\n"
    send_data += "Content-Length: "
    send_data += str(os.path.getsize('html_template.html'))
    send_data += "\r\n\r\n"
    html_file = codecs.open("html_template.html", "r", "utf-8").read()
    send_data += html_file
    return send_data


def get_file_page(filename):
    send_data = "HTTP/1.1 200 OK \r\n"
    send_data += "Content-Type: text/html; charset=utf-8\r\n"
    # send_data += "Content-Length: "
    # send_data += str(os.path.getsize('html_template2.html'))
    send_data += "\r\n\r\n"
    html_file = codecs.open("html_template2.html", "r", "utf-8").read()
    # print(html_file)
    html_file = html_file.replace('PDF_filename', filename)
    # print(html_file)
    send_data += html_file
    return send_data


def get_image(conn):
    img = open('pngwing.com.png', 'rb')
    send_data = "HTTP/1.1 200 OK \r\n"
    send_data += "Content-Type: image/png\r\n"
    send_data += "Accept-Ranges: bytes"
    send_data += "Content-Length: "
    send_data += str(os.path.getsize('pngwing.com.png'))
    send_data += "\r\n\r\n"
    img_data = img.read()
    conn.sendall(send_data.encode())
    conn.sendall(img_data)
    img.close()
    '''img = Image.open('pngwing.com.png', 'r')
    conn.sendall(Image.Image.tobytes(img))'''
    return send_data

def get_not_exists_page():
    send_data = "HTTP/1.1 404 Not Found \r\n"
    send_data += "Content-Type: text/html; charset=utf-8\r\n"
    send_data += "Content-Length: "
    send_data += str(os.path.getsize('html_not_found.html'))
    send_data += "\r\n\r\n"
    html_file = codecs.open("html_not_found.html", "r", "utf-8").read()
    send_data += html_file
    return send_data


def get_not_implemented_page():
    send_data = "HTTP/1.1 501 Not Implemented \r\n"
    send_data += "Content-Type: text/html; charset=utf-8\r\n"
    send_data += "Content-Length: "
    send_data += str(os.path.getsize('html_not_implemented.html'))
    send_data += "\r\n\r\n"
    html_file = codecs.open("html_not_implemented.html", "r", "utf-8").read()
    send_data += html_file
    return send_data

def get_other_server_error_page():
    send_data = "HTTP/1.1 500 Iternal Server Error \r\n"
    send_data += "Content-Type: text/html; charset=utf-8\r\n"
    send_data += "\r\n\r\n"
    html_file = codecs.open("html_server_error.html", "r", "utf-8").read()
    send_data += html_file
    return send_data

def check_what_request(req_data):
    if req_data['Request'] == 'GET / HTTP/1.1':  # home page
        print('home page')
        return 1
    if str(req_data['Request']).find('favicon.ico') != -1:  # favicon
        # print('favicon in stupid')
        return 2
    if req_data['Request'][0:3] != 'GET':  # not a get request
        print('not GET')
        return 501
    if req_data['Request'][0:3] == 'GET' and str(req_data['Request']).find('.png') != -1:  # image request
        print("image request")
        return 4
    if req_data['Request'][0:3] == 'GET':  # file request
        print("file request")
        return 3
    return 100


def check_if_file_exists(request_dict, all_paths_dict):
    path = str(request_dict['Request']).split(' ')[1][1:]
    path += ".pdf"
    if path in all_paths_dict.keys():
        return [True, path, all_paths_dict[path]]
    return [False, path, None]


if __name__ == "__main__":

    '''    txt_string = convert_pdf_to_txt('pdfs/sentence.pdf')
        comeback = filter_stopword(txt_string)
        print(comeback)
        hw1_utils.generate_wordcloud_to_file(comeback, 'pngwing.com.png')'''

    files_data = file_runner()
    files_list = files_data[0]
    files_paths = files_data[1]
    html_file = codecs.open("html_template.html", "w", "utf-8")
    html_file.write('<!DOCTYPE html><html lang="en"> \r\n'
                    '<head><meta charset="UTF-8"> \r\n'
                    '<title>Main Page</title> \r\n'
                    '</head> \r\n'
                    '<body> \r\n'
                    '<h1> Hello, I\'m TheFlyingDutchman!</h1> \r\n')
    for key, val in files_paths.items():
        # print(key + " : " + val)
        html_file.write('<button onclick=location.href=\'http://localhost:8888/' + key.split('.')[0] + '\' type=\"button\"> ' + val
                        + '</button> <br><br> \r\n')
    html_file.write('</body> \r\n</html> \r\n')
    html_file.close()

    while True:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.bind(('localhost', SERVER_PORT))
            s.listen(1)
            conn, addr = s.accept()
            with conn:
                # print("connection")
                # print(html_file.read())
                try:
                    recv_data = conn.recv(4096)
                except:
                    continue
                else:
                    if not recv_data:
                        continue
                    nice_data = hw1_utils.decode_http(recv_data)
                    print(nice_data)
                    req_num = check_what_request(nice_data)
                    if req_num == 1:
                        send_data = get_home_page()
                        conn.sendall(send_data.encode())
                    elif req_num == 2:
                        continue
                    elif req_num == 3:
                        is_file_exists_and_path = check_if_file_exists(nice_data, files_paths)
                        is_file_exists = is_file_exists_and_path[0]
                        if is_file_exists:
                            #print("IN PDFS!!!!")
                            file_string = convert_pdf_to_txt(is_file_exists_and_path[1])
                            words = filter_stopword(file_string)
                            try:
                                hw1_utils.generate_wordcloud_to_file(words, 'pngwing.com.png')
                            except:
                                img = Image.new("RGB", (800, 1280), (220, 255, 150))
                                img.save("pngwing.com.png", "PNG")
                            send_data = get_file_page(is_file_exists_and_path[2])
                        else:
                            #print("NOT IN PDFS!!! CYBER!!!!!")
                            send_data = get_not_exists_page()
                        conn.sendall(send_data.encode())
                    elif req_num == 501:
                        send_data = get_not_implemented_page()
                        conn.sendall(send_data.encode())
                    elif req_num == 4:
                        send_data = get_image(conn)
                        # conn.sendall(send_data.encode())
                    else:
                        send_data = get_other_server_error_page()
                        conn.sendall(send_data.encode())
                finally:
                    s.close()
