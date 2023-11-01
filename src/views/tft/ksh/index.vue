<template>
  <div class="app-container">
    <h1>Food Truck Status Dashboard</h1>
    <div class="status-filters">
      <el-checkbox-group v-model="selectedStatus">
        <el-checkbox
          v-for="(status, index) in statuses"
          :key="index"
          :label="status"
        >
          <p :style="{ color: statusColor(status) }">{{ statusMap[status] }}</p>
        </el-checkbox>
      </el-checkbox-group>
    </div>
    <div class="map-container">
      <l-map :zoom="zoom" :center="center" style="height: 100vh"> <!-- 100vh 设置为全屏高度 -->
        <l-tile-layer :url="tileLayerUrl" />
        <l-marker v-for="(truck, index) in filteredTrucks" :key="index" :lat-lng="[truck.latitude, truck.longitude]">
          <l-icon :icon-url="customIcon.iconUrl" :icon-size="customIcon.iconSize" :icon-anchor="customIcon.iconAnchor" :popup-anchor="customIcon.popupAnchor">
          </l-icon>
          <l-popup>{{ truck.name }}</l-popup>
        </l-marker>
      </l-map>
    </div>
  </div>
</template>

<script>
import { LMap, LTileLayer, LMarker, LPopup, LIcon } from "vue2-leaflet";
import "leaflet/dist/leaflet.css"; // 引入 Leaflet CSS

export default {
  components: {
    LMap,
    LTileLayer,
    LMarker,
    LPopup,
    LIcon,
  },
  data() {
    return {
      statuses: ["SUSPEND", "REQUESTED", "EXPIRED", "ISSUED", "APPROVED"],
      selectedStatus: ["SUSPEND", "REQUESTED", "EXPIRED", "ISSUED", "APPROVED"],
      statusMap: {
        "SUSPEND": "可能有优惠持续关注",
        "REQUESTED": "惊喜即将诞生",
        "EXPIRED": "已过期",
        "ISSUED": "即将开业",
        "APPROVED": "广受好评"
      },
      trucks: [],
      zoom: 13,
      center: [37.7749, -122.4194],
      tileLayerUrl: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
      customIcon: {
        iconUrl: require("@/assets/marker-icon.png"),
        iconSize: [25, 41],
        iconAnchor: [12, 41],
        popupAnchor: [1, -34],
        shadowUrl: require("@/assets/marker-shadow.png"),
        shadowSize: [41, 41],
      },
    };
  },
  created() {
  },
  computed: {
    filteredTrucks() {
      return this.trucks.filter((truck) => this.selectedStatus.includes(truck.status));
    },
  },
  mounted() {
    this.$get('getTracks', '')
    .then(response => {
      if (response.code !== 0) {
        this.$message({
          message: response.data.toString(),
          type:  'error'
        });
      } else {
        this.trucks = response.data;
      }
    })
    .catch(() => {
      this.$message({
        message: 'Failed to fetch data',
        type: 'error'
      });
    });
  },
  methods:{
    statusColor(status) {
      // 根据不同的状态返回不同的颜色
      if (status === "SUSPEND") {
        return "red";
      } else if (status === "REQUESTED") {
        return "blue";
      } else if (status === "EXPIRED") {
        return "orange";
      } else if (status === "ISSUED") {
        return "green";
      } else if (status === "APPROVED") {
        return "purple";
      }
    },
  }
};
</script>

<style>
.map-container {
  height: 100vh; /* 设置 div 的高度为全屏高度 */
}
</style>
