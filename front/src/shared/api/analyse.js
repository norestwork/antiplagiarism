const BASE_URL = 'http://localhost:3000'

export async function analyseTexts(text1, text2) {
  const res = await fetch(`${BASE_URL}/analyse`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text1, text2, mode: 'general' }),
  })

  if (!res.ok) {
    const err = await res.json().catch(() => ({ text: 'Неизвестная ошибка' }))
    throw new Error(err.text ?? 'Ошибка сервера')
  }

  return res.json()
}
