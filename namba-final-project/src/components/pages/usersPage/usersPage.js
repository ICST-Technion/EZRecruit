import React, {useRef, useLayoutEffect, useState} from 'react';
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import am4themes_animated from "@amcharts/amcharts4/themes/animated";
import {AllUsersTable} from "./allUsersTable";
import { AppContext } from "../../../lib/contextLib";
import {AllVotesTable} from "./allVotesTable";
am4core.useTheme(am4themes_animated);

export const UsersPage = () => {

    const [userData,setUserData] = useState(null);
    if (userData != null) {
        console.log(userData.name)
        console.log(userData.user_ID)
    }
    /*const chart = useRef(null);
    useLayoutEffect(() => {

        if (userData != null) {
            let x = am4core.create("chartdiv", am4charts.PieChart);

            //let answers = userData.answers;
            //let answers_counter = userData.answers_counter;
            //console.log(pollData);
            //console.log(answers);
            //console.log(answers_counter);

            /*let categoryAxis = x.xAxes.push(new am4charts.CategoryAxis());
            categoryAxis.dataFields.category = "answer";
            categoryAxis.title.text = "Answers";

            let valueAxis = x.yAxes.push(new am4charts.ValueAxis());
            valueAxis.title.text = "Votes";*/


            //let series = x.series.push(new am4charts.ColumnSeries());
            //series.name = "votes";
            //let pieSeries = x.series.push(new am4charts.PieSeries());
            //pieSeries.dataFields.value = "votes";
            //pieSeries.dataFields.category = "answer";
            //series.columns.template.tooltipText = "Series: {name}\nCategory: {categoryX}\nValue: {valueY}";
            /*
            series.columns.template.fill = am4core.color("#4cd20c"); // fill
            series.dataFields.valueY = "votes";
            series.dataFields.categoryX = "answer";*/

            /*let count = Object.keys(answers).length;
            let lst = [];
            for (let i = 0; i < count; i++){
                lst.push({
                    "answer": answers[i],
                    "votes": answers_counter[i]
                });
            }

            x.data = lst;
            console.log(x.data);

            chart.current = x;

            return () => {
                x.dispose();
                //
            };
        }
    }, [userData]);
    */

    return(
        <header className="App-header">
            <AppContext.Provider value={{userData, setUserData}}>
                <AllUsersTable/>
                {userData == null ?  (<></>)
                    : (<AllVotesTable />)}

            </AppContext.Provider>
        </header>
    )
}