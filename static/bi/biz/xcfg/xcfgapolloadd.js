define(['backbone', 'form', 'godModel'], function(Backbone, Form, GodModel) {
    env = ""
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
            this.queryObject = $.fn.getQueryObject();
            env = this.queryObject.env;
        },
        render: function() {
            var self = this;
            var form = new Form({
                el: self.$el,
                modal: this.options.modal,
                cfg: {
                    'title': '新建Apollo模板',
                    'mode': 'create',
                    'save': '/cfg/add?env=' + env+'&type=1',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'CfgName',
                        'title': '文件名',
                        'type': 'text'
                    }, {
                        'name': 'AppName',
                        'title': '应用程序名',
                        'type': 'text'
                    }, {
                        'name': 'CfgFile',
                        'title': '文件内容',
                        'type': 'textxml'
                    },{
                        'name': 'ApolloTemplate',
                        'title': '模板内容',
                        'type': 'textnginx'
                    }],
                    rules: {
                        CfgName: {
                            required: true
                        },
                        CfgFile: {
                            xml: true
                        }
                    },
                    messages: {
                        CfgName: {
                            required: "配置文件必须填写"
                        }
                    },
                    "onSuccess": function(mode, response) {
                        if (response.code == 200) {
                            history.back();
                        } else {
                            $.fn.message({
                                msg: '添加配置文件失败！'
                            });
                        }
                    },
                    "validate": function(data) {
                        var result = $.fn.validateXML(data['CfgFile'])
                        var self = this
                        if (result.error_code == 0) {
                            self.godModel = new GodModel();
                            self.godModel.url = '/cfg/exists?env=' + env;
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
                                            msg: '配置文件已经存在!'
                                        });
                                    }

                                }
                            });
                        } else {
                            $.fn.message({
                                type: 'warning',
                                title: '警告',
                                msg: result.msg
                            });
                        }
                        return false;
                    }
                }
            });
            form.render();
        }
    });
    return mainView;
});