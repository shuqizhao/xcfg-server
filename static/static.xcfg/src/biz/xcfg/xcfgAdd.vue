<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "新建配置文件",
        mode: "create",
        save:
          self.getGlobalData().ApiBaseUrl +
          "/cfg/add?env=" +
          self.$route.query.env +
          "&type=0",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "CfgName",
            title: "文件名",
            type: "text"
          },
          {
            name: "AppName",
            title: "应用程序名",
            type: "text"
          },
          {
            name: "CfgFile",
            title: "文件内容",
            type: "textxml"
          }
        ],
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
        onSuccess: function(mode, response) {
          if (response.code == 200) {
            history.back();
          } else {
            self.$message({
              message: "添加配置文件失败！",
              type: "error"
            });
          }
        },
        validate: function(data, saveData) {
          var result = self.validateXML(data["CfgFile"]);
          if (result.error_code == 0) {
            $.ajax({
              type: "POST",
              xhrFields: {
                withCredentials: true
              },
              url:
                self.getGlobalData().ApiBaseUrl +
                "/cfg/exists?env=" +
                self.$route.query.env,
              data: data,
              success: function(response) {
                if (response.code == 200 && response.data == 0) {
                  saveData(data);
                } else {
                  self.$message({
                    type: "warning",
                    messages: "配置文件已经存在!"
                  });
                }
              }
            });
          } else {
            self.$message({
              type: "warning",
              message: result.msg
            });
          }
          return false;
        }
      }
    };
  }
};
</script>