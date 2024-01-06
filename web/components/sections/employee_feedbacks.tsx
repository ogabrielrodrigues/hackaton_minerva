import { EmployeeFeedbacks as EmplFeedbacks } from '@/components/employee_feedbacks'

export function EmployeeFeedbacks() {
  return (
    <section className="flex flex-col overflow-hidden gap-2">
      <h2 className="text-3xl font-bold">Seus Feedbacks</h2>

      <EmplFeedbacks />
    </section>
  )
}
