import React,{ useState, useEffect }from "react"
import {reactFormatter, ReactTabulator} from 'react-tabulator';
import "tabulator-tables/dist/css/tabulator.min.css";
import {APIBase} from "../../../config";
import {useAppContext} from "../../../lib/contextLib";
import httpClient from "../../../httpClient";



export const AllPollsTable2 = () => {

    //const [data,setData] = useState([{id:1, name:"Dutchman", noOfRequest: 42}])
    const [data,setData] = useState([]);
    const {setPollData} = useAppContext();
    const {setValue} = useAppContext();
    const columns=[
        {
            title: "poll_ID",
            field: "poll_ID",
            headerFilter: true
        },
        {
            title: "question",
            field: "question",
            headerFilter: true
        },
        {
            title: "answers",
            field: "answers",
        },
        {
            title: "answers_counter",
            field: "answers_counter",
        },
        {
            title: "multiple_choice",
            field: "multiple_choice",
        },
        {
            title: "correct_answers",
            field: "correct_answers",
        },
        {
            title: "solution",
            field: "solution",
        },
    ]



    function getData() {
        httpClient.get(APIBase + "/all_polls_data")
            .then(res => {
                console.log(res["data"]);
                let count = Object.keys(res["data"]).length;
                let lst = [];
                for (let i = 0; i < count; i++){
                    lst.push(res["data"][i]);
                }
                // console.log(lst);
                setData(lst);
                //console.log(data)
            })
            .catch( (e) => {
                alert("error has occurred");
                //console.log(e);
            });

        /*fetch(APIBase + "/all_polls_data",{method: 'GET', mode: "cors"})
            .then(res => res.json())
            .then(data => {
                //console.log("reload");
                //console.log(data);
                let count = Object.keys(data).length;
                //console.log(count);
                let lst = [];
                for (let i = 0; i < count; i++){
                    lst.push(data[i]);
                }
                //console.log(lst);
                setData(lst);
                //console.log(data)
            })
            .catch( (e) => {
                alert("error has occurred");
                //console.log(e);
            })*/
    }

    function rowClicked(e, row){
        setValue("default");
        setPollData(row.getData());
    }


    useEffect(()=>{
        //console.log("use Effect is on");
        getData();
    },[])

    const options = {
        height: '100%',
        debugInvalidOptions: false,
        selectable: 1,
        layout: "fitColumns",
        pagination: "local",
        paginationSize: 5,
    };


    return(
            <div className="AllPollsTable">
                <ReactTabulator
                    columns={columns}
                    data={data} // here is the state of the table
                    options={options}
                    rowClick={(e,row) => rowClicked(e,row)}
                />
            </div>
    );
}