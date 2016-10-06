define(['backbone'], function (Backbone) {
    var mainView = Backbone.View.extend({
        initialize: function (options) {
            this.options = options;
            this.queryObject = $.fn.getQueryObject();
        },
        render: function () {
            var data = {};
            var src = this.queryObject.url;
            if (!src) {
                src = GodData.FrameUrl;
            }
            var html = "<iframe id='myframe' ";
            //debugger;
            //if(!this.queryObject.url)
            // {
            html += " onload = 'height = this.contentDocument.body.scrollHeight' ";
            // }
            // else{
            // 	html+=" height='100%' ";
            // }
            html += " marginwidth='0' marginheight='0' frameborder='0' scrolling='auto'  style='width:100%;overflow-x:hidden;background-color:transparent; text-align:center;' src='/" + src + "'></iframe>"
            this.$el.html(html);
            this.runResizeTask();
        },
        runResizeTask: function () {
            this.intervalInt = setInterval(this.doReSize, 500); //每隔半秒执行一次
        },
        doReSize: function () {
            var myframe = $('#myframe');
            var target = myframe[0];
            if (target) {
                target.height = target.contentDocument.body.scrollHeight;
            } else {
                clearInterval(this.intervalInt);
            }
        }
    });
    return mainView;
});