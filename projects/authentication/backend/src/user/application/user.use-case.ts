import { User } from "../domian/User";

interface UserRepository {
  findById(id: string): Promise<User>;
}

interface UserValidatePassword {
  validatePassword(password: string, user: User): Promise<[boolean] | [boolean,Error]>;
}

interface GenerateUserToken {
    generateToken(user: User): Promise<string>;
}

class UserUseCase {
  constructor(
    private readonly userRepository: UserRepository,
    private readonly userValidatePassword: UserValidatePassword,
    private readonly generateUserToken: GenerateUserToken
) {}

  async login(user: User): Promise<String> {
    const user_ = await this.userRepository.findById(user.id);
    if (!user_) {
        return "error";
    }

    const [isValid,error] = await this.userValidatePassword.validatePassword(user.password, user_);
    if (error) {
        return "error";
    }

    if (!isValid) {
        return "error";
    }

    const token = await this.generateUserToken.generateToken(user_);

    return token;
  }

  async register(user: User): Promise<String> {
    const user_ = await this.userRepository.findById(user.id);
    if (!user_) {
        return "error";
    }

    const [isValid,error] = await this.userValidatePassword.validatePassword(user.password, user_);
    if (error) {
        return "error";
    }
    
    if (!isValid) {
        return "error";
    }

    const token = await this.generateUserToken.generateToken(user_);

    return token;
  }
}

export default UserUseCase;
