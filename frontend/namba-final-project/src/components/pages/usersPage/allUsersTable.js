import React,{ useState, useEffect }from "react"
import {reactFormatter, ReactTabulator} from 'react-tabulator';
import "tabulator-tables/dist/css/tabulator.min.css";
import {APIBase} from "../../../config";
import Tabulator from "tabulator-tables";
import {useAppContext} from "../../../lib/contextLib";
import httpClient from "../../../httpClient";



export const AllUsersTable = () => {

    //const [data,setData] = useState([{id:1, name:"Dutchman", noOfRequest: 42}])
    const [data,setData] = useState([]);
    const {setUserData} = useAppContext();
    const columns=[
        {
            title: "user_ID",
            field: "user_ID",
        },
        {
            title: "name",
            field: "name",
            headerFilter: true
        },
    ]



    function getData() {
        httpClient.get(APIBase + "/all_users_data")
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

        /*fetch(APIBase + "/all_users_data",{method: 'GET', mode: "cors"})
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
        setUserData(row.getData());
        //console.log(row.getData());
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
        <div className="AllUsersTable">
            <ReactTabulator
                columns={columns}
                data={data} // here is the state of the table
                options={options}
                rowClick={(e,row) => rowClicked(e,row)}
            />
        </div>
    );
}