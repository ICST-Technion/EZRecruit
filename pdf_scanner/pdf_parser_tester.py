from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
from re import M
from unicodedata import name
from pdf_parser_core import findSpecificWord
from pdf_parser_core import findWords


path = "C:/technion/semester_8/EZRecruit/pdf_scanner/aa.pdf"
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
    str = content[0]['text']
    print("STR")
    print(str)
    print("Before searching")
    print(findSpecificWord("בוקר")(str))
    wordsList = ["בוקר","טוב"]
    print(findWords(wordsList, str))


if __name__ ==  "__main__":
    main()