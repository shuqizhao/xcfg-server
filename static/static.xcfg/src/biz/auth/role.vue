<template>
  <list :cfg="cfg"></list>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "角色管理列表",
        parentTitle:'权限管理',
        url: this.getGlobalData().ApiBaseUrl + "/role/list",
        columns: [
          {
            title: "id",
            name: "Id",
            isHide: true
          },
          {
            title: "角色名",
            name: "Name"
          },
          {
            title: "创建时间",
            name: "CreateTime"
          }
        ],
        idName: "Id",
        functions: {
          common: [
            {
              text: "添加角色",
              url: "/auth/roleadd",
              mode: "navigate"
            }
          ]
        },
        operations: [
          {
            text: "查看",
            url: "/auth/roleview"
          }
        ]
      }
    };
  }
};
</script>