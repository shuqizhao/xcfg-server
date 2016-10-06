define(['backbone', 'text!createForm.html', 'godModel'], function(Backbone, jst, GodModel) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function(options) {
            this.options = options;
            this.godModel = new GodModel();
            this.godModel.url = this.options.cfg.postUrl;
        },
        events: {
            'click .btn-commit': 'btnCommit'
        },
        render: function() {
            var self = this;
            var data = {
                'cfg': this.options.cfg
            };
            self.$el.append(self.template(data));
            var dataTableCfg = {
                onfocusout: false,
                onclick: false,
                focusInvalid: false,
                onkeyup: function(element) {
                    $(element).valid();
                },
                errorPlacement: function(error, element) {
                    //console.log(error);
                    //console.log(element);
                    element.attr("data-toggle", "tooltip");
                    element.attr("data-placement", "bottom");
                    //element.attr("title", error.text());
                    element.attr("data-original-title", error.text())
                    element.tooltip("show");
                    //error.appendTo(element.parent()); 
                },
                success: function(error, element) {
                    //console.log(error);
                    //console.log(element);
                    $(element).tooltip("destroy");
                },
                submitHandler: function(form) {
                    var data = {};
                    self.$el.find(".createFrom").find(".form-control").each(function(index) {
                        data[this.id] = $(this).val();
                    });
                    var isOk = true;
                    if (self.options.cfg.validate) {
                        isOk = self.options.cfg.validate(data);
                    }

                    if (!isOk) {
                        return;
                    }
                    self.godModel.set(data);
                    self.godModel.save({}, {
                        success: function(model, response) {
                            self.$el.find(".createFrom").modal("hide");
                            $.fn.message({
                                msg: self.options.cfg.title + '成功了！'
                            });
                            if (self.options.cfg.onSuccess) {
                                self.options.cfg.onSuccess(model, response);
                            }
                            if (self.searchDataTable) {
                                self.searchDataTable.reLoad();
                            }
                        },
                        error: function(err) {
                            console.log(err);
                        }
                    });
                }
            };
            var lastCfg = $.extend(true, dataTableCfg, self.options.cfg);
            self.$el.find(".createFrom").find("form").validate(lastCfg);

        },
        btnCommit: function() {
            this.$el.find(".createFrom").find("form").submit();
        }
    });
    return mainView;
});