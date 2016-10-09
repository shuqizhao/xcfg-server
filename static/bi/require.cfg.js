require.config({
    baseUrl: 'static/bi',
    paths: {
        //ref
        text: 'ref/text',
        jquery: 'ref/jquery-1.7.2.min',
        //jquery: 'ref/jquery-1.11.1.min',
        validate: 'ref/jquery.validate',
        underscore: 'ref/underscore-min',
        backbone: 'ref/backbone-min',
        bootstarp: 'ref/bootstrap/js/bootstrap.min',
        jDataTable: 'ref/jquery.dataTables',
        bDataTable: 'ref/dataTables.bootstrap',
        toastmessage: 'ref/toastmessage/jquery.toastmessage',
        jBootstrap: 'ref/jquery.bootstrap',
        popModal: 'ref/popModal/popModal.min',
        datetimepicker: 'ref/datetimepicker/jquery.datetimepicker',
        uploadify: 'ref/uploadify/jquery.uploadify.min',
        suggest: 'ref/suggest/bootstrap-suggest.min',
        //compents
        select2select: 'compents/select2select',
        cookie: 'compents/cookie',
        godModel: 'compents/godModel',
        tree: 'compents/tree',
        dataTable: 'compents/dataTable',
        searchDataTable: 'compents/searchDataTable',
        message: 'compents/message',
        form: 'compents/form',
        modalForm: 'compents/modalForm',
        //biz common
        backbonecfg: 'backbone.cfg',
        layout: 'biz/common/layout',
        login: 'biz/login/login',
        //biz
    },
    shim: {
        "bootstarp": ["jquery"],
        "backbone": ["jquery"],
        "jDataTable": ["jquery"],
        "bDataTable": ['jDataTable', 'backbone'],
        "jBootstrap": ["backbone"],
        "form": ["jBootstrap"],
        "toastmessage": ["jquery"],
        "message": ["jquery"],
        "select2select": ["jquery"]
    }
});
require(['backbonecfg'], function(backbonecfg) {});

GodData = {
    Index: '/bi/biz/frame/frame.js',
    FrameUrl: 'static/bi/biz/home/index.html',
    RegistBiz: [
        '/static/bi/biz/frame/frame.js',
        '/static/bi/biz/auth/user.js',
        '/static/bi/biz/auth/useradd.js',
        '/static/bi/biz/auth/userview.js',
        '/static/bi/biz/auth/usermod.js',
        '/static/bi/biz/auth/role.js',
        '/static/bi/biz/auth/roleadd.js',
        '/static/bi/biz/auth/roleview.js',
        '/static/bi/biz/auth/resource.js',
    ]
};