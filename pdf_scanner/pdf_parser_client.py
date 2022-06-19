import errno
import requests
import json
from fileinput import FileInput
import PySimpleGUI as sg
import os.path
import subprocess

# sg.theme_previewer()
sg.theme('BlueMono')    # Keep things interesting for your users

API_ENDPOINT = "http://localhost:5000/cvs"


def buildJsonForReq(filePath, wordsListString):
    jsonForReq = {}
    jsonForReq["pathToDir"] = filePath
    wordsList = wordsListString.split(',')
    wordsJson = {}
    for word in wordsList:
        wordsJson[word] = "1"
    jsonForReq["wordsList"] = wordsJson
    return jsonForReq

def printResults(outputWindow, data):
    fullString = ""
    for cv in data:
        fullString += "Name: " + cv[0] + "\n"
        fullString += "Wordcount: " + str(cv[1]) + "\n"
        fullString += "Description: " + str(cv[2]) + "\n"
        fullString += "----------------------------------------------------------------\n"
    outputWindow.update(fullString)

def listBoxCvs(listBoxWindow, data, folder):
    filesList = []
    for cv in data:
        filesList.append(cv[0])
    listBoxWindow.update(filesList)

def getLastFolder():
    try:
        file = open("chosen-folder.txt", "r")
        folderpath = file.readline()
        return folderpath
    except:
        return ""

def setLastFolder(folderpath):
    file = open("chosen-folder.txt", "w+")
    file.write(folderpath)

def getLastWordsList():
    # TODO: fix problem with reading hebrew text files
    try:
        file = open("words-list.txt", "r")
        # print("check")
        wordsList = file.readline()
        print(wordsList)
        return wordsList
    except:
        # print("SHOT")
        return ""

def setLastWordsList(wordsList):
    file = open("words-list.txt", "w+")
    file.write(wordsList)

def showListOfPdfFiles(window, folder):
    window["-FOLDER-"].update(folder)
    window["-FOLDER-"].set_tooltip(folder)
    try:
        # Get list of files in folder
        file_list = os.listdir(folder)
    except:
        file_list = []

    fnames = [
        f
        for f in file_list
        if os.path.isfile(os.path.join(folder, f))
        and f.lower().endswith((".pdf"))
    ]
    window["-FILE LIST-"].update(fnames)

def initData(window):
    window.Finalize()
    showListOfPdfFiles(window, getLastFolder())
    # window["-WordsList-"].update(getLastWordsList())
    window.refresh()

def loadingWindow():
    layout = [
        [
            sg.Text("Loading...")
        ]
    ]
    window = sg.Window(title="", layout=layout)
    return window


def dirSearchWindow():
    file_list_column = [
    [
        sg.Text("CV's Folder"),
        sg.In(size=(25, 1), enable_events=True, key="-FOLDER-"),
        sg.FolderBrowse(),
    ],
    [
        sg.Text('Words List'),
        sg.Input(size=(30, 1), enable_events=True, key="-WordsList-"),
    ],
    [
        sg.Listbox(values=[], enable_events=True, size=(40, 20), key="-FILE LIST-")
    ],
    [
        sg.Column([[sg.Button("Scan", key="-SCAN-")]], justification='center')
    ]
]

    image_viewer_column = [
        [sg.Multiline(size=(50,30), key="-OUTPUT-")]
    ]

    cv_column = [
        [sg.Listbox(values=[], select_mode="LISTBOX_SELECT_MODE_SINGLE",
         enable_events=True, size=(40, 20), key="-CV LIST-")]
    ]

    layout = [
        [
            sg.Column(file_list_column),
            sg.VSeperator(),
            sg.Column(image_viewer_column),
            sg.VSeperator(),
            sg.Column(cv_column)
        ]
    ]
    window = sg.Window(title="Pdf Parser", layout=layout)
    initData(window)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break

        if event == "-FOLDER-":
            folder = values["-FOLDER-"]
            showListOfPdfFiles(window, folder)
            setLastFolder(folder)
        elif event == "-SCAN-":
            try:
                dirPath = values["-FOLDER-"]
                wordsList = values["-WordsList-"]
                jsonForReq = buildJsonForReq(dirPath, wordsList)
                r = requests.post(url = API_ENDPOINT, json = jsonForReq, headers={'content-type': 'application/json'})
                data = json.loads(r.json())
                printResults(window["-OUTPUT-"], data)
                listBoxCvs(window["-CV LIST-"], data, values["-FOLDER-"])
            except:
                print("ERROR IN PARSING")
        elif event == "-CV LIST-":
            filepath = os.path.join(values["-FOLDER-"], values["-CV LIST-"][0])
            subprocess.Popen([filepath],shell=True)
    window.close()

def fileSearchWindow():
    file_list_column = [
    [
        sg.Text("CV's Folder"),
        sg.In(size=(25, 1), enable_events=True, key="-FOLDER-"),
        sg.FolderBrowse(),
    ],
    [
        sg.Text('Words List'),
        # sg.In(size=(25, 1), enable_events=True, key="-WordsList-"),
        sg.InputText(size=(30, 1), enable_events=True, key="-WordsList-"),
        # sg.Button("Button", key="-PRESSED-")
    ],
    [
        sg.Listbox(
            values=[], enable_events=True, size=(40, 20), key="-FILE LIST-"
        )
    ],
    [
        sg.Column([[sg.Button("Scan", key="-SCAN-")]], justification='center')
    ]
]

    image_viewer_column = [
        [sg.Text("Choose a file from list on left:")],
        [sg.Text(size=(45, 1), key="-TOUT-")],
        [sg.Multiline(size=(50,10), key="-OUTPUT-")]
    ]

    cv_column = [
        [sg.Listbox(values=[], select_mode="LISTBOX_SELECT_MODE_SINGLE",
         enable_events=True, size=(40, 20), key="-CV LIST-")]
    ]

    layout = [
        [
            sg.Column(file_list_column),
            sg.VSeperator(),
            sg.Column(image_viewer_column),
            sg.VSeperator(),
            sg.Column(cv_column)
        ]
    ]
    window = sg.Window(title="Pdf Parser", layout=layout)
    initData(window)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break

        if event == "-FOLDER-":
            folder = values["-FOLDER-"]
            showListOfPdfFiles(window, folder)
            setLastFolder(folder)
        elif event == "-FILE LIST-":  # A file was chosen from the listbox
            try:
                filename = os.path.join(
                    values["-FOLDER-"], values["-FILE LIST-"][0]
                )
                window["-TOUT-"].update(filename)
                # window["-IMAGE-"].update(filename=filename)
            except:
                print("whoops")
        elif event == "-SCAN-":
            #print(values["-WordsList-"])
            try:
                filePath = filename
                # print("filePath: " + filePath)
                wordsList = values["-WordsList-"]
                # print("wordsList: " + wordsList)
                jsonForReq = buildJsonForReq(filePath, wordsList)
                r = requests.post(url = API_ENDPOINT, json = jsonForReq, headers={'content-type': 'application/json'})
                data = json.loads(r.json())
                # print(data)
                printResults(window["-OUTPUT-"], data)
                listBoxCvs(window["-CV LIST-"], data, values["-FOLDER-"])
                # window["-OUTPUT-"].print(data)
            except:
                print("ERROR IN PARSING")
        elif event == "-CV LIST-":
            filepath = os.path.join(values["-FOLDER-"], values["-CV LIST-"][0])
            subprocess.Popen([filepath],shell=True)
            

    window.close()


def mainWindow():
    layout = [
        [
            sg.Button("Search Single File", key="openFileSearch"),
            sg.Button("Search Directory", key="openDirectorySearch")

        ]
    ]
    window = sg.Window(title="Main Window", layout=layout)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break
        elif event == "openDirectorySearch":
            dirSearchWindow()
        elif event == "openFileSearch":
            fileSearchWindow()

if __name__ == "__main__":
    mainWindow()