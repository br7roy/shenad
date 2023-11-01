const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  prdEnv: state => state.app.form.prdEnv,
  schUserName: state => state.app.form.schUserName
}
export default getters
