$(function () {
    initTitles();
});

function initTitles() {
    $.ajax({
        url : "/titles",
        type : "GET",
        dataType:'json',
        success : function (data) {
            //$("#show").text(data.titleList).append(data.titleList.length)
            for (var i = 0;i<data.titleList.length;i++){
                /*$("#show").append(data.titleList[i])*/
                var title = $("<a class=\"list-group-item\" ></a>").text(data.titleList[i].Title).attr('href',"/blog/"+data.titleList[i].Id);
                $("#help_link").before(title);
            }
        }
    })
}

$(".list-group-item").click(function () {
    $.ajax({
        url : "/blog/"+this.id,
        type : "GET",
        dataType:'json',
        success : function (data) {

        }
    })
});