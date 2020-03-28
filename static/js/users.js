$(function () {
   // initUser();
});

function initUser() {
    $("tbody").empty();
    $.ajax({
        url : "/users",
        type : "GET",
        dataType:'json',
        success : function (data) {
            /*  alert(data.resultData)*/
            //console.log(data.resultData.length);
            for (var i = 0;i<data.data.length;i++){
                var td = $("<tr><td>"+data.data[i].Id+"</td><td>"+data.data[i].Username+"</td><td style='text-align: center'><button class='btn btn-primary' onclick='delUser(\""+data.data[i].Id+"\")'>删除</button></td></tr>")
                $("#users").append(td)
                /*/!*$("#show").append(data.titleList[i])*!/
                var sel = $(" <option value=\""+data.resultData[i].Id+"\">Google</option>").text(data.resultData[i].KindName);
                $("#kind").append(sel);*/
            }
        }
    })
}

function delUser(userId) {
    console.log(userId);
    $.ajax({
        url : "/user/"+userId,
        type : "DELETE",
        dataType:'json',
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                alert("删除成功");
                initUser();
            }
        }
    })
}