/**
 * Created by olehgol on 08.01.17.
 */

(function() {
    let input_form = document.getElementById("input_form");
    let salary = document.getElementsByName("salary")[0];
    let goal = document.getElementsByName("goal")[0];
    let table = document.getElementById("my_table");
    let table_body = document.getElementById("table_body");
    let row_template = document.getElementsByClassName("row_template")[0];
    let fields = ["percent", "result_income", "time"];
    function load_table(salary, goal, cb) {
        let oReq = new XMLHttpRequest();
        oReq.addEventListener("load", cb);
        oReq.open("GET", "/data?salary=" + salary + "&goal=" + goal);
        oReq.send();
    }

    input_form.addEventListener("submit", function(e) {
        e.preventDefault();
        load_table(salary.value, goal.value, function (e) {
            let target = e.currentTarget;
            if (target.status == 200) {
                let response = JSON.parse(target.response);
                console.log(response);
                table_body.innerHTML = "";
                for (let i in response) {
                    let row = row_template.cloneNode(true);
                    row.className = "";
                    for (let field of fields){
                        let element = row.getElementsByClassName(field)[0];
                        element.innerHTML= response[i][field]
                    }
                    table_body.appendChild(row);
                }
                table.style.display=  null;
            }
        });
        return false;
    })
})();
