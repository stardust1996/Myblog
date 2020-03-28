var comments = $(" <form id='reviews_comment_form' >" +
    " <input type='hidden' name='superior' id='rSuperior'>" +
    " <input type='hidden' name='superior_user' id='rSuperiorUser'>" +
    " <input type='hidden' name='uid' id='rUid'>" +
    " <input type='hidden' name='bid' id='rBid'>" +
    "<div class=\"input-group mb-3\">\n" +
    "    <input type=\"text\" class=\"form-control\" name='content'>\n" +
    "   <div class=\"input-group-append\">\n" +
    "      <div class=\"btn btn-success\"  id='reviews_comment' onclick='reviewsComment()'>评论</div>  \n" +
    "   </div>\n" +
    "</div>" +
    "</form>");

//var uid = $("#userId").val();
var bid = window.location.href.substring(window.location.href.lastIndexOf("/")+1);

$(function () {
    initComment();
    $("#reviews").click(sendComments)
});


function initComment() {
    //console.log(url.substring(url.lastIndexOf("/")+1));
   // alert(bid);
    $.ajax({
        url : "/comments?bid="+bid,
        type : "GET",
        dataType:'json',
        success : function (data) {
            for (var i = 0;i<data.resultData.length;i++){
                var cl = $("<hr class=\"simple\" color=\"#6f5499\" />\n" +
                    "                               <div class=\"row\">\n" +
                    "                                    <div class=\"col-md-8\">\n" +
                    "                                     <p><span class='font-weight-bold'>"+data.resultData[i].comment_username+"</span></p>" +
                  //  "                                        <span class=\"badge badge-default\" style=\"color: #000000\">"+data.resultData[i].comment_username+"</span>\n" +
                    "                                    </div>\n" +
                    "                                    <div class=\"col-md-4\">\n" +
                    "                                        <span class=\"text-muted\"><small>"+data.resultData[i].comment_created.substring(0,10)+"</small></span>\n" +
                    "                                    </div>\n" +
                    "                                </div>\n" +
                    "                                <div class=\"row\">\n" +
                    "                                    <div class=\"col-md-12\">\n" +
                    "                                        <p style='padding-left: 10px;margin: 0'>\n" +
                    data.resultData[i].comment_content+"\n" +
                    "                                        </p>\n" +
                    "                                    </div>" +
                    "                                    </div>" +
                    "<button  class=\"text-muted btn\" id=\"btn"+i+"\" style='font-size: 10px' >回复</button>" +
                    "");
                if (data.resultData[i].comment_superior_user != null && data.resultData[i].comment_superior_user !== "" ){
                    //console.log(cl.children(".col-md-8").children("p"))
                    cl.children(".col-md-8").children("p").append("&nbsp; <span class='text-info'>回复</span>&nbsp; "+data.resultData[i].comment_superior_user)
                }

                $("#comments").append(cl);
               reviews("btn"+i,data.resultData[i].comment_user_id,data.resultData[i].comment_id)
            }
        }
    })
}

function reviews(btnId,sUid,cid) {
    //console.log($("#userId").val());
    var uid = $("#userId").val();
    $("#"+btnId).click(function () {
        if (uid != null && uid !== ""){
            //alert("aaaaa")
            $(this).after(comments);
            $(this).next().children("#rSuperior").val(cid);
            $(this).next().children("#rSuperiorUser").val(sUid);
            $(this).next().children("#rUid").val(uid);
            $(this).next().children("#rBid").val(bid);
            //$("#reviews_comment").click(reviewsComment)
           /* $("#reviews_comment").click(function () {
                $.ajax({
                    url: "/comment",
                    type: "POST",
                    dataType: "json",
                    data:$("#reviews_comment_form").serialize(),
                    success:function (data) {
                        alert("ssss");
                        alert(bid);
                        if (data.msg!=null &&data.msg!==""){
                            alert(data.msg)
                        }else {
                            alert("/blog/"+bid);
                            window.location.href = "/blogs";

                        }
                    }
                })
            })*/
            //console.log($(this).next().children("#superior") );
        }else {
            alert("请先登录再回复！")
        }

    })
}

function reviewsComment() {
    $.ajax({
        url: "/comment",
        type: "POST",
        async:false,
        dataType: "json",
        data:$("#reviews_comment_form").serialize(),
        success:function (data) {
            if (data.msg!=null &&data.msg!==""){
                alert(data.msg)
            }else {
                window.location.reload()

            }
        }
    });
    return false;
}

function sendComments() {
    $("#bid").val(bid);
    $("#uid").val($("#userId").val());
    $.ajax({
        url: "/comment",
        type: "POST",
        async:false,
        dataType: "json",
        data:$("#reviews_form").serialize(),
        success:function (data) {
            if (data.msg!=null &&data.msg!==""){
                alert(data.msg)
            }else {
                window.location.reload()
            }
        }
    })
}