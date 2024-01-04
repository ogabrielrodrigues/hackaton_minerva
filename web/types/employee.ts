import { Feedback } from "./feedback"

export type Employee = {
  registry: string
  name: string
  email: string
  sector: string
  unit: string
  administrator: boolean
  feedbacks?: Feedback[]
}
