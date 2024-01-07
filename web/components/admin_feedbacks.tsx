import { Feedback as FeedbackType } from '@/types/feedback'
import { api_url } from '@/config/api'
import { GetEmployee } from '@/hooks/get_employee'
import { Feedback } from './feedback'
import { GetToken } from '@/hooks/get_token'
import { FilterForm } from './forms/filter_form'
import { GetFilter } from '@/hooks/get_filter'

export async function AdminFeedbacks() {
  const { token } = await GetToken()
  const { employee } = await GetEmployee()
  const { filter } = await GetFilter()

  const url = filter
    ? `${api_url}/feedback/filter?unit=${employee?.unit}&sector=${filter}`
    : `${api_url}/feedback/filter?unit=${employee?.unit}`

  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const raw_feedbacks = (await response.json()) as FeedbackType[]
  const feedbacks = raw_feedbacks?.filter(
    feedback =>
      feedback.employee_registry != employee?.registry && !feedback.answered
  )

  return (
    <>
      <FilterForm />
      {feedbacks?.length != 0 ? (
        <ul className="flex flex-col gap-2 overflow-y-scroll last:mb-12">
          {feedbacks?.map(feedback => (
            <Feedback
              key={feedback.id}
              id={feedback.id}
              administrator={true}
              answered={feedback.answered}
              answer_id={feedback.answer_id}
              employee_registry={feedback.employee_registry}
              content={feedback.content}
              sent_at={feedback.sent_at}
            />
          ))}
        </ul>
      ) : (
        <p className="self-center mt-16 font-medium text-primary">
          Nenhum feedback disponível!
        </p>
      )}
    </>
  )
}
