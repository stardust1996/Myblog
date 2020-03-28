$(function () {
    initBlogList();
    $("#edit_blog").click(updateBlog);
});


function initBlogList() {
    $("#blog_list").empty();
    $.ajax({
        url : "/blogs?edit=1",
        type : "GET",
        dataType:'json',
        success : function (data) {
            /*  alert(data.resultData)*/
            //console.log(data.resultData.length);
            for (var i = 0;i<data.data.length;i++){
                var td = $("<tr><td>"+data.data[i].Title+"</td><td>"+data.data[i].KindName+"</td>" +
                    "<td style='text-align: center'>" +
                    "<button class='btn btn-primary ' style='margin-left: 5px;margin-right: 5px' onclick='delBlog(\""+data.data[i].Id+"\")'>删除</button>" +
                    "<button class='btn btn-primary ' style='margin-left: 5px;margin-right: 5px' onclick='getBlogById(\""+data.data[i].Id+"\")'>修改</button></td></tr>")
                $("#blog_list").append(td)
                /*/!*$("#show").append(data.titleList[i])*!/
                var sel = $(" <option value=\""+data.resultData[i].Id+"\">Google</option>").text(data.resultData[i].KindName);
                $("#kind").append(sel);*/
            }
        }
    })
}


function delBlog(bid) {
    $.ajax({
        url : "/blog/"+bid,
        type : "DELETE",
        dataType:'json',
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                alert("删除成功");
                initBlogList();
            }
        }
    })
}

function getBlogById(bid) {
    $.ajax({
        url : "/blog/"+bid+"?edit=1",
        type : "GET",
        dataType:'json',
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                $("#title_list").hide();
                $("#edit").show();
                $("#edit_Id").val(data.data.id);
                $("#edit_title").val(data.data.title);
                $("#edit_kind").val(data.data.kind);
                $("#edit_content").val(data.data.content);
            }
        }
    })
}

function updateBlog() {
    $.ajax({
        url: "/blog",
        type: "PUT",
        dataType: "json",
        data : $("#edit_blog_form").serialize(),
        success : function (data) {
            if (data.msg != null && data.msg !== ""){
                alert(data.msg)
            }else {
                alert("修改保存成功！");
                initBlogList();
                editReset();
            }
        }
    })
}

function editReset() {
    $("#title_list").show();
    $("#edit").hide();
    $("#edit_blog_form")[0].reset();
    $("#edit_Id").val("");
}