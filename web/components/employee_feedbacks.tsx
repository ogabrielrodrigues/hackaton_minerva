import { api_url } from '@/config/api'
import { GetEmployee } from '@/hooks/get_employee'
import { Feedback } from './feedback'
import { GetToken } from '@/hooks/get_token'
import { Employee } from '@/types/employee'

export async function EmployeeFeedbacks() {
  const { token } = await GetToken()
  const { employee } = await GetEmployee()

  const response = await fetch(`${api_url}/employee/${employee?.registry}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const feedbacks = ((await response.json()) as Employee).feedbacks

  return (
    <>
      {feedbacks?.length != 0 ? (
        <ul className="flex flex-col gap-2 overflow-y-scroll last:mb-12">
          {feedbacks?.map(feedback => (
            <Feedback
              key={feedback.id}
              id={feedback.id}
              answered={feedback.answered}
              administrator={false}
              answer_id={feedback.answer_id}
              employee_registry={feedback.employee_registry}
              content={feedback.content}
              sent_at={feedback.sent_at}
            />
          ))}
        </ul>
      ) : (
        <p className="self-center mt-16 font-medium text-primary">
          Você ainda não possui feedbacks !
        </p>
      )}
    </>
  )
}
