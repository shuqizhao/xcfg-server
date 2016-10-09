define(['backbone', 'text!modalForm.html', 'form', 'popModal'], function(Backbone, jst, Form) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function(options) {
            this.options = options;
        },
        render: function() {
            var self = this;
            var data = {};
            var appendAaer = self.$el.find('.aModalForm');
            if (appendAaer.length <= 0) {
                self.$el.prepend("<div class='modal fade aModalForm' tabindex='-1' role='dialog' aria-labelledby='myModalLabel' style='width:630px;'></div>")
            }
            self.$el.find('.aModalForm').html(self.template(data));
            var Module = require('/static/bi/biz/' + this.options.url + '.js')
            var m = new Module({
                el: self.$el,
                modal: true
            });
            m.render();
        }
    });
    return mainView;
});