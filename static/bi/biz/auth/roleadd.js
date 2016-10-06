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
                    'title': '添加角色',
                    'mode': 'create',
                    'save': '/role/add',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'Name',
                        'title': '角色名',
                        'type': 'text'
                    }],
                    rules: {
                        Name: {
                            required: true
                        }
                    },
                    messages: {
                        Name: {
                            required: "角色名必须填写"
                        }
                    },
                    "onSuccess": function(mode, response) {
                        if (response.code == 200) {
                            history.back();
                        } else {
                            $.fn.message({
                                msg: '添加角色失败！'
                            });
                        }
                    },
                    "validate": function(data) {
                        self = this
                        this.godModel = new GodModel();
                        self.godModel.url = '/role/exists';
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
                                        msg: '角色已经存在!'
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