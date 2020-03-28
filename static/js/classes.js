$(function () {
   // initClasses();
    $("#addClass").unbind("click").click(function () {
        addClasses();
    })
});

function initClasses() {
    $("tbody").empty();
    $.ajax({
        url : "/classes",
        type : "GET",
        dataType:'json',
        success : function (data) {
            /*  alert(data.resultData)*/
            //console.log(data.resultData.length);
            for (var i = 0;i<data.resultData.length;i++){
                var td = $("<tr><td>"+data.resultData[i].Id+"</td><td>"+data.resultData[i].KindName+"</td><td style=\"text-align: center\"><button class='btn btn-primary' onclick='delClass("+data.resultData[i].Id+")'>删除</button></td></tr>")
                $("#classes").append(td)
                /*/!*$("#show").append(data.titleList[i])*!/
                var sel = $(" <option value=\""+data.resultData[i].Id+"\">Google</option>").text(data.resultData[i].KindName);
                $("#kind").append(sel);*/
            }
        }
    })
}

function delClass(classId) {
    $.ajax({
        url : "/classes/"+classId,
        type : "DELETE",
        dataType:'json',
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                alert("删除成功");
                initClasses();
                initSelClasses();
            }
        }
    })
}

function addClasses() {
    $.ajax({
        url : "/class",
        type : "POST",
        dataType:'json',
        data: {kind_name : $("#kind_name").val()},
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                alert("添加成功");
                $("#kind_name").val("");
                initClasses();
                initSelClasses();
            }
        }
    })
}