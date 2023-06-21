export function getQueryParam(url: string, key: string): string | null {
  const index = url.indexOf('?')
  if (index === -1) {
    return null
  }

  const params = url.substring(index + 1).split('&')
  for (const param of params) {
    const [paramKey, paramValue] = param.split('=')
    if (paramKey === key) {
      return paramValue
    }
  }
  return null
}
