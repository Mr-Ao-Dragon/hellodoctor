var {spawn} = require('child_process');
//换成你自己的hbuilderX path
// const {cli} = require('C:/HBuilder/HBuilderX/cli.exe')
(function(){
  return new Promise(async (resolve, reject) => {
    let resk = await spawn('C:/HBuilder/HBuilderX/cli.exe', ['publish','--platform','H5','--project','hellodoctor'])
    resk.on('close', async (code) => {
      console.log('app资源本地打包完成')
      resolve('success')
    })
  })
})()
