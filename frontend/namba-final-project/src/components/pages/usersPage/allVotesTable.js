import React,{ useState, useEffect }from "react"
import {reactFormatter, ReactTabulator} from 'react-tabulator';
import "tabulator-tables/dist/css/tabulator.min.css";
import {APIBase} from "../../../config";
import Tabulator from "tabulator-tables";
import {useAppContext} from "../../../lib/contextLib";
import {UsersPage} from "./usersPage";
import httpClient from "../../../httpClient";



export const AllVotesTable = () => {
    //const [data,setData] = useState([{id:1, name:"Dutchman", noOfRequest: 42}])
    const [data,setData] = useState([]);
    const {userData} = useAppContext();
    const columns=[

        {
            title: "poll_ID",
            field: "poll_ID",
            width: 200,
        },
        {
            title: "question",
            field: "question",
            width: 350,
        },
        {
            title: "answers",
            field: "answers",
            headerSort:false,
            width: 350,
        },
    ]



    function getData() {
        httpClient.get(APIBase + "/specific_user_answers/" + userData.user_ID)
            .then(res => {
                //console.log(res["data"]);
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


        /*fetch(APIBase + "/specific_user_answers/" + userData.user_ID,{method: 'GET', mode: "cors"})
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
        //alert(row.getData());
        //console.log(row.getData());
        //setUserData(row.getData());
        //console.log(row.getData());
    }


    useEffect(()=>{
        //console.log("use Effect is on");
        getData();
    },[userData])

    const options = {
        height: '100%',
        debugInvalidOptions: false,
        selectable: 1,
        layout: "fitColumns",
        pagination: "local",
        paginationSize: 5,
    };


    return(
        <div className="AllVotesTable">
            <ReactTabulator
                columns={columns}
                data={data} // here is the state of the table
                options={options}
                rowClick={(e,row) => rowClicked(e,row)}
            />
        </div>
    );
}