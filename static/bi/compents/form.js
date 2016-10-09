define(['backbone', 'text!form.html', 'godModel', 'bootstarp', 'popModal', 'datetimepicker', 'uploadify', 'suggest'], function(Backbone, jst, GodModel) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function(options) {
            this.options = options;
            this.godModel = new GodModel();
        },
        events: {
            'click .btn-commit': 'btnCommit',
            'click .btn-cancel': 'btnCancel',
            'click .btn-edit': 'btnEdit',
            'click .btn-decommit': 'btnDecommit',
            'click .btn-decancel': 'btnDecancel',
            'click .btn-buttons': 'btnButtons'
        },
        render: function() {
            var self = this;
            if (self.options.cfg.mode == 'detailEdit' && self.options.cfg.detailEditMode != 'edit') {
                self.options.cfg.title = self.options.cfg.detailTitle;
            } else if (self.options.cfg.mode == 'detailEdit' && self.options.cfg.detailEditMode == 'edit') {
                self.options.cfg.title = self.options.cfg.editTitle;
            }
            if (self.options.cfg.mode != 'create') {
                self.godModel.url = this.options.cfg.get.url;
                self.godModel.fetch({
                    data: this.options.cfg.get.params,
                    success: function(model, text) {
                        var detail = model.toJSON();
                        var title = self.options.cfg.title;
                        if (title) {
                            var items = self.options.cfg.items;
                            for (var i = 0; i < items.length; i++) {
                                if (title.indexOf(items[i].name) > 0) {
                                    title = title.replace('{' + items[i].name + '}', detail.data[items[i].name]);
                                }
                            }
                            self.options.cfg.title = title;
                        }
                        var data = {
                            'cfg': self.options.cfg,
                            'detail': detail.data
                        };
                        if (self.options.modal) {
                            self.$el.find('.modal-body').html(self.template(data));
                        } else {
                            self.$el.html(self.template(data));
                        }
                        if (self.options.cfg.afterEditRender) {
                            self.options.cfg.afterEditRender();
                        }
                    },
                    error: function() {}
                });
            } else {
                var data = {
                    'cfg': self.options.cfg,
                    'detail': {}
                };
                if (self.options.modal) {
                    self.$el.find('.modal-body').html(self.template(data));
                } else {
                    self.$el.html(self.template(data));
                }
            }
            if (self.options.modal) {
                if (self.$el.find(".aModalForm")) {
                    //self.$el.find(".aModalForm").find("form")[0].reset();
                    self.$el.find(".aModalForm").find("[data-toggle='tooltip']").tooltip("destroy");
                    self.$el.find(".aModalForm").modal("show");
                }
            }
        },
        btnCommit: function(e, handler) {
            var self = this;
            var validateCfg = {
                onfocusout: false,
                onclick: false,
                focusInvalid: false,
                onkeyup: function(element) {
                    //console.log(element);
                    $(element).valid();
                },
                errorPlacement: function(error, element) {
                    if (element.attr('controltype') == 'suggest') {
                        element.next().find('button').attr("data-toggle", "tooltip");
                        element.next().find('button').parent().attr("data-placement", "right");
                        element.next().find('button').parent().attr("data-original-title", error.text());
                        element.next().find('button').parent().tooltip("show");
                    } else if (element.attr('controltype') == 'upload') {
                        element.parent().find('object').attr("data-toggle", "tooltip");
                        element.parent().find('object').parent().attr("data-placement", "right");
                        element.parent().find('object').parent().attr("data-original-title", error.text());
                        element.parent().find('object').parent().tooltip("show");
                    } else {
                        element.attr("data-toggle", "tooltip");
                        element.attr("data-placement", "right");
                        element.attr("data-original-title", error.text());
                        element.tooltip("show");
                    }
                    //error.appendTo(element.parent()); 
                },
                success: function(error, element) {
                    //console.log(error);
                    //console.log(element);
                    $(element).tooltip("destroy");
                },
                submitHandler: function(form) {
                    var data = self.getData()
                        //console.log(data);
                    var isOk = true;
                    if (self.options.cfg.validate) {
                        //自定义验证
                        isOk = self.options.cfg.validate(data);
                    }

                    if (!isOk) {
                        return;
                    }
                    if (self.options.cfg.beforeCommit) {
                        self.options.cfg.beforeCommit(data);
                    }
                    self.saveData(data, handler)
                }
            };
            var lastCfg = $.extend(true, validateCfg, self.options.cfg);
            self.$el.find("form").validate(lastCfg);
        },
        btnCancel: function() {
            var self = this;
            if (self.options.cfg.onCancel) {
                self.options.cfg.onCancel();
            } else if (self.options.modal) {
                //$('.modal-backdrop').hide();
            } else {
                //$.fn.navigate();
                history.go(-1);
            }
        },
        btnEdit: function() {
            this.options.cfg.detailEditMode = 'edit';
            this.render();
        },
        btnDecommit: function() {
            self = this;
            self.btnCommit(null, function() {
                self.options.cfg.detailEditMode = 'detail';
                self.render();
            });

        },
        btnDecancel: function() {
            this.options.cfg.detailEditMode = 'detail';
            this.render();
        },
        btnButtons: function(e) {
            var name = e.target.name;
            var buttons = this.options.cfg.buttons;
            for (var i = 0; i < buttons.length; i++) {
                if (buttons[i].name == name) {
                    buttons[i].handler();
                }
            };
        },
        getData: function() {
            var self = this
            var data = {};
            self.$el.find(".form").find(".form-control").each(function(index) {
                var item = $(this);
                if (this.type == 'checkbox') {
                    if (item.attr('checked')) {
                        data[this.id] = 1;
                    } else {
                        data[this.id] = 0;
                    }
                } else if (item.attr('controltype') == 'suggest') {
                    data[this.id] = item.attr('data-id');
                } else if (item.attr('controltype') == 'addline') {
                    var arraytemp = [];
                    item.find('input').each(function() {
                        var input = $(this);
                        arraytemp.push(input.val());
                    });
                    data[this.id] = arraytemp;
                } else {
                    data[this.id] = item.val();
                }
            });
            return data;
        },
        saveData: function(data, handler) {
            self = this;
            if (data == null) {
                data = self.getData()
            }
            var saveGodModel = new GodModel();
            saveGodModel.url = self.options.cfg.save;
            saveGodModel.set(data);
            saveGodModel.save({}, {
                success: function(model, response) {
                    if (response.code && response.code == "201") {
                        window.open(response.data);
                    } else if (response.code && response.code == "203") {
                        alert("查看console.log");
                        console.log(response.data);
                    }
                    if (response.code == "200") {
                        $.fn.message({
                            msg: self.options.cfg.title + '成功了！'
                        });
                        //如果是模态窗体,关闭模态
                        $('.aModalForm').modal('hide');
                    } else if (response.message) {
                        $.fn.message({
                            type: 'warning',
                            title: '警告',
                            msg: response.message
                        });
                    }
                    if (self.options.cfg.onSuccess) {
                        if (self.options.cfg.onSuccess(model, response)) {
                            if (handler) {
                                handler();
                            } else {
                                $.fn.navigate();
                            }
                        }
                    } else {
                        if (handler) {
                            handler();
                        } else {
                            $.fn.navigate();
                        }
                    }
                },
                error: function(err) {
                    console.log(err);
                }
            });
        }
    });
    return mainView;
});