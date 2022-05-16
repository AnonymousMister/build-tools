CHAT_CHANNEL 10.6 所有 触发ChatOps命令的 源聊天通道 CHAT_INPUT 10.6 所有 在ChatOps命令中 传递的其他参数 CI 所有 0.4 标记作业在CI环境中执行 CI_API_V4_URL 11.7
所有 GitLab API v4根URL CI_BUILDS_DIR 所有 11.10 执行构建的顶级目录。 CI_COMMIT_BEFORE_SHA 11.2 所有
先前的最新提交存在于分支中。始终0000000000000000000000000000000000000000处于合并请求的管道中。 CI_COMMIT_DESCRIPTION 10.8 所有
提交的描述：如果标题少于100个字符，则不带第一行的消息；在其他情况下为完整消息。 CI_COMMIT_MESSAGE 10.8 所有 完整的提交消息。 CI_COMMIT_REF_NAME 9.0 所有 构建项目的分支或标记名称
CI_COMMIT_REF_PROTECTED 11.11 所有 true如果作业在受保护的引用上运行，false则不是 CI_COMMIT_REF_SLUG 9.0 所有
$CI_COMMIT_REF_NAME小写，缩短为63个字节，并与一切除了0-9和a-z与更换-。没有前导/尾随-。在URL，主机名和域名中使用。 CI_COMMIT_SHA 9.0 所有 为其构建项目的提交修订
CI_COMMIT_SHORT_SHA 11.7 所有 的前八个字符 CI_COMMIT_SHA CI_COMMIT_BRANCH 12.6 0.5 提交分支名称。存在于分支管道中，包括默认分支的管道。在合并请求管道中不存在。
CI_COMMIT_TAG 9.0 0.5 提交标记名称。仅在构建标签时显示。 CI_COMMIT_TITLE 10.8 所有 提交的标题-消息的第一行 CI_COMMIT_TIMESTAMP 13.4 所有 ISO
8601格式的提交时间戳记。 CI_CONCURRENT_ID 所有 11.10 单个执行程序中生成执行的唯一ID。 CI_CONCURRENT_PROJECT_ID 所有 11.10 单个执行者和项目中的构建执行的唯一ID。
CI_CONFIG_PATH 9.4 0.5 CI配置文件的路径。默认为.gitlab-ci.yml CI_DEBUG_TRACE 所有 1.7 是否启用 调试日志记录（跟踪） CI_DEFAULT_BRANCH 12.4 所有
项目默认分支的名称。 CI_DEPLOY_FREEZE 13.2 所有 true如果管道在部署冻结窗口期间运行，则包含在值中。 CI_DEPLOY_PASSWORD 10.8 所有 GitLab
Deploy令牌的身份验证密码，仅在项目具有相关性时才提供。 CI_DEPLOY_USER 10.8 所有 GitLab Deploy令牌的身份验证用户名，仅在项目具有相关性时才存在。 CI_DISPOSABLE_ENVIRONMENT
所有 10.1 标记该作业是在一次性环境中执行的（仅为该作业创建并在执行后处置/销毁的事物-shell和以外的所有执行者ssh）。如果环境是一次性的，则将其设置为true，否则将完全未定义。 CI_ENVIRONMENT_NAME 8.15
所有 该作业的环境名称。仅在environment:name设置时存在。 CI_ENVIRONMENT_SLUG 8.15 所有 环境名称的简化版本，适用于包含在DNS，URL，Kubernetes标签等中。仅在environment:
name设置时存在。 CI_ENVIRONMENT_URL 9.3 所有 此作业的环境的URL。仅在environment:url设置时存在。 CI_EXTERNAL_PULL_REQUEST_IID 12.3 所有
如果管道用于外部请求，则来自GitHub的请求请求ID 。仅在使用only: [external_pull_requests]或rules语法且拉取请求处于打开状态时可用。
CI_EXTERNAL_PULL_REQUEST_SOURCE_REPOSITORY 13.3 所有 如果管道用于外部请求请求，则请求请求的源存储库名称。仅在使用only: [external_pull_requests]
或rules语法且拉取请求处于打开状态时可用。 CI_EXTERNAL_PULL_REQUEST_TARGET_REPOSITORY 13.3 所有
如果管道用于外部请求，则请求请求的目标存储库名称。仅在使用only: [external_pull_requests]或rules语法且拉取请求处于打开状态时可用。
CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_NAME 12.3 所有 如果管道用于外部请求，则请求请求的源分支名称。仅在使用only: [external_pull_requests]
或rules语法且拉取请求处于打开状态时可用。 CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_SHA 12.3 所有 如果管道用于外部请求，则请求请求的源分支的HEAD SHA
。仅在使用only: [external_pull_requests]或rules语法且拉取请求处于打开状态时可用。 CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_NAME 12.3 所有
如果管道用于外部请求，则请求请求的目标分支名称。仅在使用only: [external_pull_requests]或rules语法且拉取请求处于打开状态时可用。
CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_SHA 12.3 所有 如果管道用于外部请求，则请求请求目标分支的HEAD SHA 。仅在使用only: [external_pull_requests]
或rules语法且拉取请求处于打开状态时可用。 CI_HAS_OPEN_REQUIREMENTS 13.1 所有 true仅当管道的项目有任何开放需求时才包含在值中。如果管道项目没有开放要求，则不包括在内。 CI_JOB_ID 9.0 所有
GitLab CI / CD在内部使用的当前作业的唯一ID CI_JOB_IMAGE 12.9 12.9 运行CI作业的图像的名称 CI_JOB_MANUAL 8.12 所有 指示作业已手动启动的标志 CI_JOB_NAME 9.0 0.5
在中定义的作业名称 .gitlab-ci.yml CI_JOB_STAGE 9.0 0.5 在中定义的阶段名称 .gitlab-ci.yml CI_JOB_STATUS 所有 13.5
每个跑步者阶段的工作状态都被执行。与使用after_script，其中CI_JOB_STATUS可以是：success，failed或canceled。 CI_JOB_TOKEN 9.0 1.2
令牌，用于通过一些API端点进行身份验证并下载相关的存储库。只要作业正在运行，令牌就有效。 CI_JOB_JWT 12.10 所有 RS256 JSON
Web令牌，可用于与支持JWT身份验证的第三方系统进行身份验证，例如HashiCorp的Vault。 CI_JOB_URL 11.1 0.5 职位详情网址 CI_KUBERNETES_ACTIVE 13.0 所有
true仅当管道具有可用于部署的Kubernetes群集时，才包含在值中。如果没有群集，则不包括在内。可作为替代only:kubernetes/except:kubernetes与rules:if
CI_MERGE_REQUEST_ASSIGNEES 11.9 所有 如果管道用于合并请求，则合并请求的受让人用户名的逗号分隔列表。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。
CI_MERGE_REQUEST_ID 11.6 所有 合并请求的实例级别ID。仅当管道用于合并请求并且创建合并请求时才可用。这是GitLab上所有项目的唯一ID。 CI_MERGE_REQUEST_IID 11.6 所有
合并请求的项目级IID（内部ID）。仅当管道用于合并请求并且创建合并请求时才可用。该ID对于当前项目是唯一的。 CI_MERGE_REQUEST_LABELS 11.9 所有
如果管道用于合并请求，则合并请求的逗号分隔标签名。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_MILESTONE 11.9 所有
如果管道用于合并请求，则合并请求的里程碑标题。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_PROJECT_ID 11.6 所有
如果管道用于合并请求，则合并请求的项目的ID 。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_PROJECT_PATH 11.6 所有
如果管道用于合并请求（例如namespace/awesome-project），则合并请求的项目路径。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。
CI_MERGE_REQUEST_PROJECT_URL 11.6 所有 如果管道用于合并请求（例如http://192.168.10.15:3000/namespace/awesome-project），则合并请求项目的URL
。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_REF_PATH 11.6 所有
如果管道用于合并请求，则合并请求的ref路径。（例如refs/merge-requests/1/head）。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。
CI_MERGE_REQUEST_SOURCE_BRANCH_NAME 11.6 所有 如果管道用于合并请求，则合并请求的源分支名称。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。
CI_MERGE_REQUEST_SOURCE_BRANCH_SHA 11.9 所有 如果管道用于合并请求，则合并请求的源分支的HEAD SHA 。仅在使用only: [merge_requests]
或rules语法，创建合并请求且管道为合并结果管道时可用。 CI_MERGE_REQUEST_SOURCE_PROJECT_ID 11.6 所有 如果管道用于合并请求，则合并请求的源项目的ID
。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_SOURCE_PROJECT_PATH 11.6 所有
如果管道用于合并请求，则合并请求的源项目的路径。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_SOURCE_PROJECT_URL 11.6 所有
如果管道用于合并请求，则合并请求的源项目的URL 。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_TARGET_BRANCH_NAME 11.6 所有
如果管道用于合并请求，则合并请求的目标分支名称。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_TARGET_BRANCH_SHA 11.9 所有
如果管道用于合并请求，则合并请求的目标分支的HEAD SHA 。仅在使用only: [merge_requests]或rules语法，创建合并请求且管道为合并结果管道时可用。 CI_MERGE_REQUEST_TITLE 11.9 所有
如果管道用于合并请求，则合并请求的标题。仅当使用only: [merge_requests]或rules语法并且创建合并请求时可用。 CI_MERGE_REQUEST_EVENT_TYPE 12.3 所有
合并请求的事件类型（如果管道用于合并请求）。可能是detached，merged_result或merge_train。 CI_NODE_INDEX 11.5 所有 作业在作业集中的索引。如果作业未并行化，则不会设置此变量。
CI_NODE_TOTAL 11.5 所有 并行运行的此作业的实例总数。如果作业未并行化，则此变量设置为1。 CI_PAGES_DOMAIN 11.8 所有 托管GitLab页面的已配置域。 CI_PAGES_URL 11.8 所有
GitLab页面构建页面的URL。始终属于的子域CI_PAGES_DOMAIN。 CI_PIPELINE_ID 8.10 所有 当前管道的实例级别ID。这是GitLab上所有项目的唯一ID。 CI_PIPELINE_IID 11.0 所有
当前管道的项目级IID（内部ID）。该ID对于当前项目是唯一的。 CI_PIPELINE_SOURCE 10.0 所有
指示如何触发管道。可能的选项有：push，web，schedule，api，external，chat，webide，merge_request_event，external_pull_request_event，parent_pipeline，trigger，或pipeline（更名为cross_project_pipeline自13.0）。对于在GitLab
9.5之前创建的管道，此显示为unknown。 CI_PIPELINE_TRIGGERED 所有 所有 指示已触发作业的标志 CI_PIPELINE_URL 11.1 0.5 管道详细资料网址 CI_PROJECT_DIR 所有 所有
克隆存储库以及运行作业的完整路径。如果设置了GitLab Runnerbuilds_dir参数，则相对于的值设置此变量builds_dir。有关更多信息，请参见GitLab Runner的高级配置。 CI_PROJECT_ID 所有 所有
GitLab CI / CD在内部使用的当前项目的唯一ID CI_PROJECT_NAME 8.10 0.5
当前正在构建的项目的目录名称。例如，如果项目的URL是gitlab.example.com/group-name/project-1，该CI_PROJECT_NAME会project-1。 CI_PROJECT_NAMESPACE 8.10
0.5 当前正在构建的项目名称空间（用户名或组名） CI_PROJECT_ROOT_NAMESPACE 13.2 0.5
当前正在构建的根项目名称空间（用户名或组名）。例如，如果CI_PROJECT_NAMESPACE是root-group/child-group/grandchild-group，CI_PROJECT_ROOT_NAMESPACE是root-group。
CI_PROJECT_PATH 8.10 0.5 具有项目名称的名称空间 CI_PROJECT_PATH_SLUG 9.3 所有 $CI_PROJECT_PATH小写并与除一切0-9，并a-z代之以-。在URL和域名中使用。
CI_PROJECT_REPOSITORY_LANGUAGES 12.3 所有 信息库中使用的语言的逗号分隔小写列表（例如ruby,javascript,html,css） CI_PROJECT_TITLE 12.4 所有
可读的项目名称，显示在GitLab Web界面中。 CI_PROJECT_URL 8.10 0.5 访问项目的HTTP（S）地址 CI_PROJECT_VISIBILITY 10.3 所有 项目可见性（内部，私人，公共）
CI_REGISTRY 8.10 0.5 如果启用了Container Registry，它将返回GitLab的Container Registry的地址。:port如果在注册表配置中指定了一个变量，则此变量包括一个值。
CI_REGISTRY_IMAGE 8.10 0.5 如果为项目启用了容器注册表，则它返回绑定到特定项目的注册表地址 CI_REGISTRY_PASSWORD 9.0 所有 用于将容器推送到当前项目的GitLab容器注册表的密码。
CI_REGISTRY_USER 9.0 所有 用于将容器推送到当前项目的GitLab容器注册表的用户名。 CI_REPOSITORY_URL 9.0 所有 克隆Git存储库的URL CI_RUNNER_DESCRIPTION 8.10
0.5 保存在GitLab中的跑步者的描述 CI_RUNNER_EXECUTABLE_ARCH 所有 10.6 GitLab Runner可执行文件的操作系统/体系结构（请注意，它不一定与执行程序的环境相同） CI_RUNNER_ID
8.10 0.5 正在使用的跑步者的唯一ID CI_RUNNER_REVISION 所有 10.6 正在执行当前作业的GitLab Runner版本 CI_RUNNER_SHORT_TOKEN 所有 12.3
跑步者令牌的前八个字符用于验证新的作业请求。用作跑步者的唯一ID CI_RUNNER_TAGS 8.10 0.5 定义的运行器标签 CI_RUNNER_VERSION 所有 10.6 正在执行当前作业的GitLab Runner版本
CI_SERVER 所有 所有 标记作业在CI环境中执行 CI_SERVER_URL 12.7 所有 GitLab实例的基本URL，包括协议和端口（如https://gitlab.example.com:8080）
CI_SERVER_HOST 12.1 所有 GitLab实例URL的主机组件，不带协议和端口（如gitlab.example.com） CI_SERVER_PORT 12.8 所有
GitLab实例URL的端口组件，不包含主机和协议（例如3000） CI_SERVER_PROTOCOL 12.8 所有 GitLab实例URL的协议组件，不带主机和端口（例如https） CI_SERVER_NAME 所有 所有
用于协调作业的CI服务器的名称 CI_SERVER_REVISION 所有 所有 用于计划作业的GitLab修订版 CI_SERVER_VERSION 所有 所有 用于计划作业的GitLab版本
CI_SERVER_VERSION_MAJOR 11.4 所有 GitLab版本主要组件 CI_SERVER_VERSION_MINOR 11.4 所有 GitLab版本次要组件 CI_SERVER_VERSION_PATCH 11.4
所有 GitLab版本补丁组件 CI_SHARED_ENVIRONMENT 所有 10.1
标记作业是在共享环境中执行的（在CI调用（例如executorshell或sshexecutor）中持续存在的内容）。如果共享环境，则将其设置为true，否则将完全未定义。 GITLAB_CI 所有 所有 标记作业在GitLab CI /
CD环境中执行 GITLAB_FEATURES 10.6 所有 以逗号分隔的实例和计划可用的许可功能列表 GITLAB_USER_EMAIL 8.12 所有 开始工作的用户的电子邮件 GITLAB_USER_ID 8.12 所有
开始工作的用户的ID GITLAB_USER_LOGIN 10.0 所有 开始工作的用户的登录用户名 GITLAB_USER_NAME 10.0 所有 开始工作的用户的真实姓名