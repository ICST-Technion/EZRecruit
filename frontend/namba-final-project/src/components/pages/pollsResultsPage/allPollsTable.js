import React,{ useState, useEffect }from "react"
import {reactFormatter, ReactTabulator} from 'react-tabulator';
import "tabulator-tables/dist/css/tabulator.min.css";
import {APIBase} from "../../../config";
import {useAppContext} from "../../../lib/contextLib";
import httpClient from "../../../httpClient";



export const AllPollsTable = () => {

    //const [data,setData] = useState([{id:1, name:"Dutchman", noOfRequest: 42}])
    const [data,setData] = useState([]);
    const {setPollData} = useAppContext();
    const columns=[
        {
            title: "jobTitle",
            field: "title",
            headerFilter: true,
            width: 175
        },
        {
            title: "jobDescription",
            field: "description",
            headerFilter: true
        },
        {
            title: "jobLocation",
            field: "location",
            headerSort:false,
        },
        /*{
            title: "labels",
            field: "answers_counter",
            headerSort:false,
        },
        {
            title: "jobId",
            field: "multiple_choice",
        },*/
    ]


//var url = 'https://ezrecruit-backend-ryo2vcvbqq-uc.a.run.app/' + 'jobs'
    function getData() {
        httpClient.get("https://ezrecruit-backend-ryo2vcvbqq-uc.a.run.app/jobs")
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