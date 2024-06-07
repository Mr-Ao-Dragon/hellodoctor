<script>
	import request from "@/api/request";
	let searchParams = new URLSearchParams(window.location.search)
	export default {
		data() {
			return {
				ready: true,
			}
		},
		async onLoad(options) {
			await this.getData();
			const res = await request.login(searchParams.get("code"))
			uni.setStorageSync('access_token', res.access_token)
			uni.setStorageSync('expries_in', res.expries_in)
		},
	}
</script>

<style lang="scss">

</style>