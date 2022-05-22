from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
import pathlib
from re import M
from pdf_parser_core import *


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


if __name__ ==  "__main__":
    # main()
    testDirSearch()
