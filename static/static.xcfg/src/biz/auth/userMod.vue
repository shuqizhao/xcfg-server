<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "修改密码",
        mode: "create",
        save: self.getGlobalData().ApiBaseUrl + "/auth/mod",
        modal: "true",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "OldPwd",
            title: "原密码",
            type: "pwd"
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
        validate: function(data, saveData) {
          $.ajax({
            type: "POST",
            xhrFields: {
              withCredentials: true
            },
            url: self.getGlobalData().ApiBaseUrl + "/auth/checkpwd",
            data: data,
            success: function(response) {
              if (response.code == 200 && response.data == 1) {
                saveData(data);
              } else {
                self.$message({
                  type: "warning",
                  message: "原密码输入错误!"
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