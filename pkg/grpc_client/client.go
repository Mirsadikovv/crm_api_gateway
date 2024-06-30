package grpc_client

import (
	"crm_api_gateway/config"
	"crm_api_gateway/genproto/administrator_service"
	"crm_api_gateway/genproto/branch_service"
	"crm_api_gateway/genproto/event_registrate_service"
	"crm_api_gateway/genproto/event_service"
	"crm_api_gateway/genproto/group_service"
	"crm_api_gateway/genproto/journal_service"
	"crm_api_gateway/genproto/lesson_service"
	"crm_api_gateway/genproto/manager_service"
	"crm_api_gateway/genproto/perfomance_service"
	"crm_api_gateway/genproto/schedule_service"
	"crm_api_gateway/genproto/student_service"
	"crm_api_gateway/genproto/superadmin_service"
	"crm_api_gateway/genproto/support_teacher_service"
	"crm_api_gateway/genproto/teacher_service"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	Student() student_service.StudentServiceClient
	Teacher() teacher_service.TeacherServiceClient
	SupportTeacher() support_teacher_service.SupportTeacherServiceClient
	Manager() manager_service.ManagerServiceClient
	Branch() branch_service.BranchServiceClient
	Group() group_service.GroupServiceClient
	Administrator() administrator_service.AdministratorServiceClient
	Superadmin() superadmin_service.SuperadminServiceClient
	Event() event_service.EventServiceClient
	EventRegistrate() event_registrate_service.EventRegistrateServiceClient
	Lesson() lesson_service.LessonServiceClient
	Journal() journal_service.JournalServiceClient
	Schedule() schedule_service.ScheduleServiceClient
	Perfomance() perfomance_service.PerfomanceServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connUser, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %v port:%v err: %v",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connSchedule, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("schedule service dial host: %v port:%v err: %v",
			cfg.ScheduleServiceHost, cfg.ScheduleServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"student_service":          student_service.NewStudentServiceClient(connUser),
			"teacher_service":          teacher_service.NewTeacherServiceClient(connUser),
			"support_teacher_service":  support_teacher_service.NewSupportTeacherServiceClient(connUser),
			"branch_service":           branch_service.NewBranchServiceClient(connUser),
			"group_service":            group_service.NewGroupServiceClient(connUser),
			"manager_service":          manager_service.NewManagerServiceClient(connUser),
			"administrator_service":    administrator_service.NewAdministratorServiceClient(connUser),
			"superadmin_service":       superadmin_service.NewSuperadminServiceClient(connUser),
			"event_service":            event_service.NewEventServiceClient(connUser),
			"event_registrate_service": event_registrate_service.NewEventRegistrateServiceClient(connUser),
			"lesson_service":           lesson_service.NewLessonServiceClient(connSchedule),
			"schedule_service":         schedule_service.NewScheduleServiceClient(connSchedule),
			"journal_service":          journal_service.NewJournalServiceClient(connSchedule),
			"perfomance_service":       perfomance_service.NewPerfomanceServiceClient(connSchedule),
		},
	}, nil
}

func (g *GrpcClient) StudentService() student_service.StudentServiceClient {
	return g.connections["student_service"].(student_service.StudentServiceClient)
}

func (g *GrpcClient) TeacherService() teacher_service.TeacherServiceClient {
	return g.connections["teacher_service"].(teacher_service.TeacherServiceClient)
}

func (g *GrpcClient) BranchService() branch_service.BranchServiceClient {
	return g.connections["branch_service"].(branch_service.BranchServiceClient)
}

func (g *GrpcClient) GroupService() group_service.GroupServiceClient {
	return g.connections["group_service"].(group_service.GroupServiceClient)
}

func (g *GrpcClient) SupportTeacherService() support_teacher_service.SupportTeacherServiceClient {
	return g.connections["support_teacher_service"].(support_teacher_service.SupportTeacherServiceClient)
}

func (g *GrpcClient) AdministratorService() administrator_service.AdministratorServiceClient {
	return g.connections["administrator_service"].(administrator_service.AdministratorServiceClient)
}

func (g *GrpcClient) ManagerService() manager_service.ManagerServiceClient {
	return g.connections["manager_service"].(manager_service.ManagerServiceClient)
}

func (g *GrpcClient) EventService() event_service.EventServiceClient {
	return g.connections["event_service"].(event_service.EventServiceClient)
}

func (g *GrpcClient) EventRegistrateService() event_registrate_service.EventRegistrateServiceClient {
	return g.connections["event_registrate_service"].(event_registrate_service.EventRegistrateServiceClient)
}

func (g *GrpcClient) SuperadminService() superadmin_service.SuperadminServiceClient {
	return g.connections["superadmin_service"].(superadmin_service.SuperadminServiceClient)
}

func (g *GrpcClient) ScheduleService() schedule_service.ScheduleServiceClient {
	return g.connections["schedule_service"].(schedule_service.ScheduleServiceClient)
}

func (g *GrpcClient) LessonService() lesson_service.LessonServiceClient {
	return g.connections["lesson_service"].(lesson_service.LessonServiceClient)
}

func (g *GrpcClient) JournalService() journal_service.JournalServiceClient {
	return g.connections["journal_service"].(journal_service.JournalServiceClient)
}

func (g *GrpcClient) PerfomanceService() perfomance_service.PerfomanceServiceClient {
	return g.connections["perfomance_service"].(perfomance_service.PerfomanceServiceClient)
}
