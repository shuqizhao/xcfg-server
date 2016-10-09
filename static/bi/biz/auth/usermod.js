define(['backbone', 'form', 'godModel'], function(Backbone, Form, GodModel) {
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
        },
        render: function() {
            var self = this;
            var form = new Form({
                el: self.$el,
                modal: self.options.modal,
                cfg: {
                    'title': '修改密码',
                    'mode': 'create',
                    'save': '/auth/mod',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'OldPwd',
                        'title': '原密码',
                        'type': 'pwd'
                    }, {
                        'name': 'Pwd',
                        'title': '密码',
                        'type': 'pwd'
                    }, {
                        'name': 'Pwd1',
                        'title': '确认密码',
                        'type': 'pwd'
                    }],
                    rules: {
                        OldPwd: {
                            required: true
                        },
                        Pwd: {
                            required: true
                        },
                        Pwd1: {
                            required: true,
                            equalTo: "#Pwd"
                        }
                    },
                    messages: {
                        OldPwd: {
                            required: "原密码必须填写"
                        },
                        Pwd: {
                            required: "密码必须填写"
                        },
                        Pwd1: {
                            required: "确认密码必须填写",
                            equalTo: "确认密码必须与密码相同"
                        }
                    },
                    "validate": function(data) {
                        self = this
                        self.godModel = new GodModel();
                        self.godModel.url = '/auth/checkpwd';
                        self.godModel.set(data);
                        self.godModel.save({}, {
                            success: function(model, text) {
                                var jsonResult = model.toJSON();
                                if (jsonResult.code == 200 && jsonResult.data == 1) {
                                    form.saveData(data, function() {})
                                } else {
                                    $.fn.message({
                                        type: 'warning',
                                        title: '警告',
                                        msg: '原密码输入错误!'
                                    });
                                }

                            }
                        });
                        return false;
                    }
                }
            });
            form.render();
        }
    });
    return mainView;
});