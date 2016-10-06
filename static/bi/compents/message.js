(function($) {
    $.fn.extend({
        "message": function(options) {
            var options = $.extend({
                type: 'success',
                title: '提示',
                msg: '操作成功！',
                position: 'top-center',
                inEffectDuration: 1000,
                stayTime: 3000
            }, options);
            $().toastmessage('showToast', {
                text: options.msg,
                sticky: false,
                type: options.type,
                position: options.position,
                inEffectDuration: options.inEffectDuration,
                stayTime: options.stayTime,
            });
        }
    });
    return this;
})(jQuery);