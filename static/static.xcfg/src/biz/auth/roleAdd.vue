<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "添加角色",
        mode: "create",
        save: this.getGlobalData().ApiBaseUrl + "/role/add",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "Name",
            title: "角色名",
            type: "text"
          }
        ],
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
        onSuccess: function(mode, response) {
          if (response.code == 200) {
            history.back();
          } else {
            $.fn.message({
              msg: "添加角色失败！"
            });
          }
        },
        validate: function(data, saveData) {
          $.ajax({
            type: "POST",
            xhrFields: {
              withCredentials: true
            },
            url: self.getGlobalData().ApiBaseUrl + "/role/exists",
            data: data,
            success: function(response) {
              if (response.code == 200 && response.data == 0) {
                saveData(data);
              } else {
                self.$message({
                  type: "warning",
                  message: "角色已经存在!"
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