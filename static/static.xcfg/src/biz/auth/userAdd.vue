<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "添加用户",
        mode: "create",
        save: this.getGlobalData().ApiBaseUrl + "/user/add",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "Name",
            title: "用户名",
            type: "text"
          },
          {
            name: "Pwd",
            title: "密码",
            type: "pwd"
          },
          {
            name: "Pwd1",
            title: "确认密码",
            type: "pwd"
          }
        ],
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
        onSuccess: function(mode, response) {
          if (response.code == 200) {
            history.back();
          } else {
            $.fn.message({
              msg: "添加用户失败！"
            });
          }
        },
        validate: function(data,saveData) {
          $.ajax({
            type: "POST",
            xhrFields: {
              withCredentials: true
            },
            url: self.getGlobalData().ApiBaseUrl + "/user/exists",
            data: data,
            success: function(response) {
              if (response.code == 200 && response.data == 0) {
                  saveData(data);
              } else {
                self.$message({
                  type: "warning",
                  message: "用户名已经存在!"
                });
              }
            }
          });
          return false;
        }
      }
    };
  }
};
</script>