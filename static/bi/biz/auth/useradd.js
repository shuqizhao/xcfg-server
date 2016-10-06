define(['backbone', 'form', 'godModel'], function(Backbone, Form, GodModel) {
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
        },
        render: function() {
            var self = this;
            var form = new Form({
                el: self.$el,
                modal: this.options.modal,
                cfg: {
                    'title': '添加用户',
                    'mode': 'create',
                    'save': '/user/add',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'Name',
                        'title': '用户名',
                        'type': 'text'
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
                        Name: {
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
                        Name: {
                            required: "用户名必须填写"
                        },
                        Pwd: {
                            required: "密码必须填写"
                        },
                        Pwd1: {
                            required: "确认密码必须填写",
                            equalTo: "确认密码必须与密码相同"
                        }
                    },
                    "onSuccess": function(mode, response) {
                        if (response.code == 200) {
                            history.back();
                        } else {
                            $.fn.message({
                                msg: '添加用户失败！'
                            });
                        }
                    },
                    "validate": function(data) {
                        self = this
                        self.godModel = new GodModel();
                        self.godModel.url = '/user/exists';
                        self.godModel.set(data);
                        self.godModel.save({}, {
                            success: function(model, text) {
                                var jsonResult = model.toJSON();
                                if (jsonResult.code == 200 && jsonResult.data == 0) {
                                    form.saveData(data)
                                } else {
                                    $.fn.message({
                                        type: 'warning',
                                        title: '警告',
                                        msg: '用户名已经存在!'
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