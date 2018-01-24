<template>
  <list :cfg="cfg"></list>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "资源管理列表",
        parentTitle:'权限管理',
        url: this.getGlobalData().ApiBaseUrl + "/menu/list",
        columns: [
          {
            title: "id",
            name: "Id",
            isHide: true
          },
          {
            title: "资源名",
            name: "Name"
          },
          {
            title: "URL",
            name: "Url"
          }
        ],
        idName: "Id"
      }
    };
  }
};
</script>