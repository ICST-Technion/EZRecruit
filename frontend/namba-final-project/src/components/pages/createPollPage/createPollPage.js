import React, {useState} from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "./createPollPage.css"
import {APIBase} from "../../../config";
import {AppContext, useAppContext} from "../../../lib/contextLib";
import {AllPollsTable2} from "./allPollsTable2";
import httpClient from "../../../httpClient";


const AnswersToChooseFrom = (pollData) =>{
    const {value, setValue} = useAppContext();
    const {valueKey, setValueKey} = useAppContext();
    let answers = pollData.value.answers;
    let checks = [];
    let count = answers.length;
    for (let i = 0; i < count; i++) {
        checks.push(<option key={i} data-key={i} value={answers[i]}> {answers[i]} </option>);
    }


    function handleChange(chosen_answer){
        //console.log("change");
        //console.log(chosen_answer);
        setValue(chosen_answer.value);
        //console.log(chosen_answer.value);
        //console.log(chosen_answer.options.selectedIndex-1);
        setValueKey(chosen_answer.options.selectedIndex-1);
    }

    return(
        <>
            <h4> Choose Answer</h4>
            <select value={value} onChange={(e) => handleChange(e.target)}>
                <option key={100} value="default" disabled>
                    Choose Answer
                </option>
                {checks}
            </select>
        </>
    );
}


const CreatePollPage = () => {
    //const {isAuthenticated} = useAppContext();
    //console.log(isAuthenticated);
    const [question, setQuestion] = useState("");
    const [option1, setoption1] = useState("");
    const [option2, setoption2] = useState("");
    const [option3, setoption3] = useState("");
    const [option4, setoption4] = useState("");
    const [multiple, setMultiple] = useState(false);
    const [toWho, setToWho] = useState("all");
    const [pollData,setPollData] = useState(null);
    const [value,setValue] = useState("default");
    const [valueKey,setValueKey] = useState(null);
    function validateForm() {
    return question.length > 0 && option1.length > 0 && option2.length > 0
        && ( toWho === "all" || (toWho!=="all" && value !== "default"));
  }

  function sendToWho(sendies) {
        setToWho(sendies);
        if (sendies === "all"){

        }
        else{
            setPollData(null);
            //alert("not all");
        }
      //console.log(sendies);
  }

    function handleSubmit(event){
        event.preventDefault();
        //alert("Poll submitted!!")
        // init_poll: question, answers[]
        let answers = [option1, option2, option3, option4];
        let question2 = question.replaceAll("?","%3F");
        //alert(multiple);
        //let data = {'question': question, 'answers': answers};
        if (toWho === "all") {
            httpClient.post(APIBase + "/send_poll_to_all/" + question2 + "/" + answers + "/" + multiple)
                .then(res => {
                    if (res["data"].result === true) {
                        alert("submitted successfully")
                    }
                    if (res["data"].result === "empty") {
                        alert("No users in the system")
                    }
                })
                .catch(e => {
                    console.log(e);
                    alert("error has occurred");
                });
            /*fetch(APIBase + "/send_poll_to_all/" + question2 + "/" + answers + "/" + multiple, {
                method: 'POST',
                mode: "cors",
            })
                .then(res => res.json())
                .then(data => {
                    //console.log(data);
                    if (data.result === true) {
                        alert("submitted successfully")
                    }
                    if (data.result === "empty") {
                        alert("No users in the system")
                    }
                })
                .catch(e => {
                    console.log(e);
                    alert("error has occurred");
                });*/
        }
        else {
            let poll_id = pollData.poll_ID;
            //console.log(row);
            console.log(poll_id);
            console.log(valueKey);
            console.log(answers);
            httpClient.post(APIBase + "/send_to_specific_voters/" + poll_id + "/" + valueKey
                + "/" + question2 + "/" + answers
                + "/" + multiple)
                .then(res => {
                    if (res["data"].result === true) {
                        alert("submitted successfully")
                    }
                    if (res["data"].result === "empty") {
                        alert("No such users in the system")
                    }
                })
            .catch(e => {
                    console.log(e);
                    alert("error has occurred");
                });


            /*fetch(APIBase + "/send_to_specific_voters/" + poll_id + "/" + valueKey
                + "/" + question + "/" + answers
                + "/" + multiple, {
                method: 'POST',
                mode: "cors",
            })
                .then(res => res.json())
                .then(data => {
                    //console.log(data);
                    if (data.result === true) {
                        alert("submitted successfully")
                    }
                    if (data.result === "empty") {
                        alert("No such users in the system")
                    }
                })
                .catch(e => {
                    console.log(e);
                    alert("error has occurred");
                });*/
        }

    }

    return(
        <header className="App-header">
        <div className={"createPollPage"}>
            <h1>Create Poll Page</h1>
            <Form onSubmit={handleSubmit}>
                <Form.Group size="lg" controlId="pollQuestion">
                    <Form.Label>Poll Question</Form.Label>
                    <Form.Control placeholder="Enter poll question (required)" value={question}
                                  as="textarea" rows="3"
                                  onChange={(e) => setQuestion(e.target.value)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="pollOption1">
                    <Form.Control placeholder="Option 1 (required)" value={option1}
                                  onChange={(e) => setoption1(e.target.value)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="pollOption2">
                    <Form.Control placeholder="Option 2 (required)" value={option2}
                                  onChange={(e) => setoption2(e.target.value)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="pollOption3">
                    <Form.Control placeholder="Option 3 (optional)" value={option3}
                                  onChange={(e) => setoption3(e.target.value)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="pollOption4">
                    <Form.Control placeholder="Option 4 (optional)" value={option4}
                                  onChange={(e) => setoption4(e.target.value)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="multiple">
                    <Form.Check label="multiple choice" checked={multiple} onChange={e => setMultiple(e.target.checked)}/>
                </Form.Group>
                <Form.Group size="lg" controlId="sendToWho">
                    <select value={toWho} onChange={(e) => sendToWho(e.target.value)}>
                        <option value="all"> ALL </option>
                        <option value="by poll"> by Poll results </option>
                    </select>
                </Form.Group>
                    <AppContext.Provider value={{pollData, setPollData, value, setValue, valueKey, setValueKey}}>
                        {toWho === "all" ? (<></>)
                            : (<AllPollsTable2/>)}
                        {toWho === "all" || pollData === null ? (<></>)
                        : (<AnswersToChooseFrom value={pollData}/>)}
                    </AppContext.Provider>
                <div className={"longShot"}>
                    <Button className="custom-btn" type="submit" block size="lg" disabled={!validateForm()}
                            alignment="center"> Submit
                        Poll</Button>
                </div>
            </Form>
        </div>
            </header>
    )
}

export default CreatePollPage;