from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
import pathlib
from re import M
from pdf_parser_core import *
import tkinter as tk
from tkinter import filedialog, simpledialog
import requests
import codecs
import json


path = "C:/technion/semester_8/EZRecruit/pdf_scanner/2-pages-pdf.pdf"
lng = "heb"


def main():
    print("start")
    pdf_document = Document(
        document_path=path,
        language=lng
        )
    print(pdf_document)
    pdf2text = PDF2Text(document=pdf_document)
    content = pdf2text.extract()
    print(content)
    str = concatPagesStrings(content)
    print("STR")
    print(str)
    print("Before searching")
    print(findSpecificWord("בכלל")(str))
    wordsList = ["זה"]
    print(findWords(wordsList, str))


def testDirSearch():
    wordsList = ["עמוד","בוקר","טוב"]
    path = pathlib.Path(__file__).parent.absolute()
    print("path = ", path)
    results = searchCVsInFolder(path,wordsList)
    print(results)


def fileExplorerTest():
    root = tk.Tk()
    root.withdraw()
    # dir_path = "C:/technion/semester_8/EZRecruit/pdf_scanner"
    dir_path = filedialog.askdirectory()
    # print(file_path)

    # enter words to search for
    words = simpledialog.askstring("input", "blah", parent=root)
    # print(words)
    wordsList = words.split(',')
    # print(wordsList)

    jsonForReq = {}
    jsonForReq["pathToDir"] = dir_path
    wordsJson = {}
    for word in wordsList:
        wordsJson[word] = "1"
    jsonForReq["wordsList"] = wordsJson
    print(jsonForReq)

    API_ENDPOINT = "http://localhost:5000/cvs"
    """jsonForReq = {
	"pathToDir": "C:/technion/semester_8/EZRecruit/pdf_scanner",
	"wordsList": {"בוקר": "1", "עמוד":"1"}
}"""

    r = requests.post(url = API_ENDPOINT, json = jsonForReq, headers={'content-type': 'application/json'})
    # print(r.content)
    # print(r.json())

    data = json.loads(r.json())
    print(data)
    print(type(data))

    # got the response, now printing to a txt file
    # TODO: output file in the same directory
    o_path = "C:/technion/semester_8/EZRecruit/pdf_scanner/text_outputs/out.txt"
    with codecs.open(o_path, "w", "utf-8") as o:
        for cv in data:
            print(cv)
            o.write("Name: " + cv[0] + "\n")
            o.write("Wordcount: " + str(cv[1]) + "\n")
            o.write("Description: " + str(cv[2]) + "\n")
            o.write("----------------------------------------------------------------\n")
    

def testMapping():
    filePath = "C:/technion/semester_8/EZRecruit/pdf_scanner/daniel.pdf"
    wordMapping = getMappingOfWordsInCV(filePath)
    o_path = "C:/technion/semester_8/EZRecruit/pdf_scanner/text_outputs/mapping.txt"
    with codecs.open(o_path, "w", "utf-8") as o:
        o.write(str(wordMapping))


if __name__ ==  "__main__":
    # main()
    # testDirSearch()
    # fileExplorerTest()
    testMapping()
