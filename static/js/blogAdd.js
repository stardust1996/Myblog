$(function () {
    initSelClasses();
    $("#create_blog").click(create_blog)
});


function initSelClasses() {
    $("select").empty();
    $.ajax({
        url : "/classes",
        type : "GET",
        dataType:'json',
        success : function (data) {
          /*  alert(data.resultData)*/
            console.log(data.resultData.length);
            for (var i = 0;i<data.resultData.length;i++){
                /*$("#show").append(data.titleList[i])*/
                var sel = $(" <option value=\""+data.resultData[i].Id+"\">Google</option>").text(data.resultData[i].KindName);
                $("select").append(sel);
            }
        }
    })
}

function create_blog() {
    // console.log($("#blog_form").serialize());
    //alert("砸瓦鲁多"+$("#blog_form").serialize());
    $.ajax({
        url: "/blog",
        type: "POST",
        dataType: "json",
        data : $("#blog_form").serialize(),
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                window.location.href = "/blogs"
                /*$().confirm({
                    title: "提示",
                    content: '保存成功!是否退回到主页',
                    buttons: {
                        confirm:function () {
                            window.location.href = "/blogs"
                        },
                        cancel:function () {

                        }
                    }
                    })*/
               // alert("保存成功!");

            }
        }
    })
}