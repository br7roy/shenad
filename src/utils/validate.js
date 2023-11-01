/**
 * Created by PanJiaChen on 16/11/18.
 */

/**
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validUsername(str) {
  return /^\w+$/.test(str)
}

export function validSchUsername(str) {
  console.log('kaishijiaoyan')
  console.log(str)
  return /[a-zA-Z]/.test(str)
}
