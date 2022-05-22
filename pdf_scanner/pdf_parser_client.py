import PySimpleGUI as sg
import os.path

# sg.theme_previewer()
sg.theme('BlueMono')    # Keep things interesting for your users


def dirSearchWindow():
    file_list_column = [
    [
        sg.Text("CV's Folder"),
        sg.In(size=(25, 1), enable_events=True, key="-FOLDER-"),
        sg.FolderBrowse(),
    ],
    [
        sg.Listbox(
            values=[], enable_events=True, size=(40, 20), key="-FILE LIST-"
        )
    ],
]

    image_viewer_column = [
        [sg.Text("Choose an image from list on left:")],
        [sg.Text(size=(40, 1), key="-TOUT-")],
        [sg.Image(key="-IMAGE-")],
    ]


    layout = [
        [
            sg.Column(file_list_column),
            sg.VSeperator(),
            sg.Column(image_viewer_column)
        ]
    ]
    window = sg.Window(title="Pdf Parser", layout=layout)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break

        if event == "-FOLDER-":
            folder = values["-FOLDER-"]
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
        elif event == "-FILE LIST-":  # A file was chosen from the listbox
            pass
            """try:
                filename = os.path.join(
                    values["-FOLDER-"], values["-FILE LIST-"][0]
                )
                window["-TOUT-"].update(filename)
                window["-IMAGE-"].update(filename=filename)
            except:
                print("whoops")
                pass"""

    window.close()


def mainWindow():
    layout = [
        [
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

if __name__ == "__main__":
    mainWindow()