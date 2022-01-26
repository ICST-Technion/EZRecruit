from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
import logging
import re
logging.basicConfig(level=logging.INFO)
import sys
import codecs

path = "C:/Users/aradk/Desktop/Yearly_Project/aa.pdf"
search_for = "איש"
lng = "heb"

def findWholeWord(w):
    return re.compile(r'\b({0})\b'.format(w), flags=re.IGNORECASE).search

def get_path_from_user(path):
    print("Do you want to insert a path to read from? y/n")
    ans = input()
    if ans == "y":
        print("Enter path now: ")
        path = input()
    elif ans == "n":
        print("Using path: ", path)
    else:
        print("Bad input! Fuck you Mirotic!!!")

def get_language_and_word(lng, search_for):
    print("Enter language (3 letters)")
    lng = input()
    print("Enter sentence to search for")
    search_for = input()


def output_pdf(path, lng, search=False):
    '''search = False
    #get_path_from_user(path)
    #get_language_and_word(lng, search_for)
    if len(sys.argv) == 3:
        search = False
    elif len(sys.argv) == 4:
        search = True
    else:
        print("Wrong number of arguments")
        exit(-1)

    path = sys.argv[1]
    lng = sys.argv[2]
    if search == True:
        search_for = sys.argv[3]
    '''
    #print("args are: " + path + "  " + lng)

    ## create document for extraction with configurations
    pdf_document = Document(
        document_path=path,
        language=lng
        )
    pdf2text = PDF2Text(document=pdf_document)
    content = pdf2text.extract()
    str = content[0]['text']

    if search == False:
        f = codecs.open("demo.txt", "w", "utf-8")
        #print(str)
        f.write(str)
        f.close()
        return str

    print("\n\n\n*************************************\n")
    if (findWholeWord(search_for)(str)) != None:
        print("PDF contains the word")
    else:
        print("PDF does not contain the word")
    print("\n*************************************\n\n")
    return str

"""
    Gets 2 arguments- a pdf path and a language (usually 'heb' or 'eng') and returns
    the words from this language in this pdf as a string.
"""
def main():
    search = False
    #get_path_from_user(path)
    #get_language_and_word(lng, search_for)
    if len(sys.argv) == 3:
        search = False
    elif len(sys.argv) == 4:
        search = True
    else:
        print("Wrong number of arguments")
        exit(-1)

    path = sys.argv[1]
    lng = sys.argv[2]
    if search == True:
        search_for = sys.argv[3]

    #print("args are: " + path + "  " + lng)

    ## create document for extraction with configurations
    pdf_document = Document(
        document_path=path,
        language=lng
        )
    pdf2text = PDF2Text(document=pdf_document)
    content = pdf2text.extract()
    str = content[0]['text']

    if search == False:
        f = codecs.open("demo.txt", "w", "utf-8")
        print(str)
        f.write(str)
        f.close()
        return str

    print("\n\n\n*************************************\n")
    if (findWholeWord(search_for)(str)) != None:
        print("PDF contains the word")
    else:
        print("PDF does not contain the word")
    print("\n*************************************\n\n")
    return str
    #print(content[0]['text'])

if __name__ == "__main__":
    main()