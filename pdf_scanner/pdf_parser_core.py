from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
import re
import os

"""
    # TODO: core functions
    Single CV:
    - check if a specific (single) word is in CV
    - check if multiple words are in CV - return number of words found

    Directory:
    - check all CV's in the directory for words
"""


"""
    Tries to find a word in a CV:
    @param word: the word to search
    @return: regex function
"""
def findSpecificWord(word):
    return re.compile(r'\b({0})\b'.format(word), flags=re.IGNORECASE).search


"""
    Find a batch of words in a given CV
    @param wordsList [List[string]]: List of words to search for
    @param cv [string]: CV to search 
    @return hitsCount [int]: Number of hits
"""
def findWords(wordsList, cv):
    hitsCount = 0
    for word in wordsList:
        result = findSpecificWord(word)(cv)
        if (result != None):
            hitsCount += 1
    return hitsCount


def searchCVsInFolder(folderPath, wordsList):
    # iterate over the Cv's in folder
    for filename in os.listdir(folderPath):
        f = os.path.join(folderPath, filename)
        # checking if it's a file (a CV)
        if os.path.isfile(f):
            hitsCount = findWords

