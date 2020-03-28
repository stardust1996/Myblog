

$(function () {
    //导航栏切换
    $(".nav,.flex-column,.nav-pills").children().click(function () {
        $(this).siblings().children("a").removeClass("active");
        $(this).children("a").addClass("active");
        // console.log($(this).index())
        $($("#page").children().get($(this).index())).show().siblings().hide();
        if ($(this).index() === 0){
            editReset();
            initBlogList();
        }else if ($(this).index() === 1){
            $("#blog_form")[0].reset();
        } else if ($(this).index() === 2){
            initUser();
        }else if ($(this).index() === 3){
            initClasses();
        }
    })
});