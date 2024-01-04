export type JwtAuthPayload = {
  registry: string
  email: string
  administrator: boolean
  exp: number
}
