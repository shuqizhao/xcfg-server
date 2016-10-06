define(['backbone', 'text!login.html', 'godModel'], function(Backbone, jst, GodModel) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function() {
            this.godModel = new GodModel();
        },
        events: {
            "click #btnLogin": "login"
        },
        render: function() {
            document.title = "远程配置文件管理系统";
            var data = {};
            this.$el.html(this.template(data));
            $("#validatetip").html("");
        },
        beforeLogin: function() {
            var username = $("#un").val();
            if (username == "") {
                $("#validatetip").html("用户名不能为空!");
                return false;
            }
            var password = $("#pwd").val();
            if (password == "") {
                $("#validatetip").html("密码不能为空!");
                return false;
            }
            return true;
        },
        login: function() {
            var self = this;
            if (self.beforeLogin()) {
                self.godModel.url = '/auth/login';
                self.godModel.save({
                    un: $("#un").val(),
                    pwd: $("#pwd").val()
                }, {
                    success: function(model, text) {
                        var result = model.toJSON();
                        if (result.code == 200 && result.data == 1) {
                            location.reload()
                        } else {
                            $("#validatetip").html("用户名或密码错误!");
                        }
                    },
                    error: function() {}
                });
            }
        }
    });
    return mainView;
});