$(function () {
    //initClasses();
    UserMe();
});

function search() {
    alert("假装搜一下")
}

function initClasses() {
    $.ajax({
        url : "/classes",
        type : "GET",
        dataType:'json',
        success : function (data) {
          /*  alert(data.resultData)*/
            for (var i = 0;i<data.resultData.length;i++){
                /*$("#show").append(data.titleList[i])*/
                var cl = $("<a class=\"dropdown-item\" ></a>").text(data.resultData[i].KindName).attr('href',"/blog?kind="+data.resultData[i].Id);
                $("#classes-box").append(cl);
            }
        }
    })
}

function UserMe() {
    $.ajax({
        url : "/user/me",
        type : "GET",
        dataType:'json',
        success : function (data) {
            if (data.data != null){
                $("#sign_in").parent().hide();
                $("#user_info").show();
                $("#username").text(data.data.username);

                $("#userId").val(data.data.id);
                //console.log(data.data);
                if (data.data.id.substring(0,2) === '01'){
                    $("#admin_page").show()
                }
            }
        }
    })
}

function logout() {
    $.ajax({
        url : "/logout",
        type : "DELETE",
        dataType:'json',
        success : function (data) {
            if (data.code != null && data.code !== 0){
                alert(data.msg)
            }else {
                $("#sign_in").parent().show();
                $("#user_info").hide();
                $("#userId").val('');
                $("#username").text('');
            }
        }
    })
}

function admin_page() {

}