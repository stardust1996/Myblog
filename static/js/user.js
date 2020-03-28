$(function () {

});

//登录
function signIn() {
    $.ajax({
        url : "/user",
        type : "GET",
        dataType:'json',
        data:$('#user_form').serialize(),
        success : function (data) {
            //alert(data.msg);
            if (data.msg !== null && data.msg !==''){
                $("#alert_box").show();
                $("#alert").html("<strong>"+data.code+"</strong> "+data.msg)
            }else {
                window.location.href = "/blogs"
            }
        }
    })
}
//注册
function signUp() {
    $.ajax({
        url : "/user",
        type : "POST",
        dataType:'json',
        data:$('#user_form').serialize(),
        success : function (data) {
            if (data.msg !== null && data.msg !==''){
                $("#alert_box").show();
                $("#alert").html("<strong>"+data.Code+"</strong> "+data.msg)
            }else {
                window.location.href = "/blogs"
            }
        }
    })
}
