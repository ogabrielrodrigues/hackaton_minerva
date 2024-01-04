import { Feedbacks } from '../feedbacks'

export function AdministratorFeedbacks() {
  return (
    <section className="flex flex-col overflow-hidden gap-2">
      <h2 className="text-3xl font-bold">Feedbacks</h2>

      <Feedbacks />
    </section>
  )
}
