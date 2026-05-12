export function getScoreColor(value) {
  if (value <= 0.2) return 'var(--green)'
  if (value <= 0.5) return 'var(--yellow)'
  return 'var(--red)'
}

export function getScoreBg(value) {
  if (value <= 0.2) return 'var(--green-dim)'
  if (value <= 0.5) return 'var(--yellow-dim)'
  return 'var(--red-dim)'
}

export function getVerdict(value) {
  if (value <= 0.2) return 'Уникальный текст'
  if (value <= 0.5) return 'Частичное совпадение'
  return 'Высокое сходство'
}

export function toPct(value) {
  return Math.round((value ?? 0) * 100) + '%'
}
