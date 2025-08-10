
import { Model, CreationAttributes, ModelStatic, WhereOptions } from 'sequelize';
import { NotFoundError } from '../utils/errors/app.error';

// here we are creating a abstract class BaseRepository that will be used as a base class for all repositories like HotelRepository, RoomRepository, etc.
// if we use abstract class then no one can create an instance of this class i.e new BaseRepository() will not work directly.
// to create an instance of this class, we need to extend this class in another class like HotelRepository, RoomRepository, etc.
abstract class BaseRepository<T extends Model> {
    // here T is a generic type that extends Model, which means it can be any Sequelize model, like Hotel, Room, roomCategory etc.
    // ModelStatic<T> is a type that represents the model class for T, allowing you to perform operations like findByPk, findAll, create, etc. on the model, like we do in the hotel.repository.ts file on the Hotel model.
    // here if you want to constructor(model : T) instead of constructor(model: ModelStatic<T>), you will get an error because T is not a model class, it is a generic type that extends Model. So you need to use ModelStatic<T> to represent the model class for T.
    //  here in constructor(model : T) meaning of T is a model instance i.e object of the model but the property like findByPk, findAll, create, etc. are not available on the model instance, they are available on the model class. So you need to use ModelStatic<T> to represent the model class for T.
    protected model: ModelStatic<T>;
    // here following is the meaning of protected : 
    // "Protected members are accessible within the class itself and by subclasses, but not from outside the class hierarchy."
    // so here protected model: ModelStatic<T> means that this.model can be accessed only
    // "Only this class (BaseRepository) and its child classes can access the model property. It cannot be accessed from outside."

    constructor(model: ModelStatic<T>) {
        this.model = model;
    }

    async findById(id: number) : Promise<T | null> {
        const record = await this.model.findByPk(id);
        if (!record) {
            throw new NotFoundError(`Record with id ${id} not found`);
        }
        return record;
    }

    async findAll(): Promise<T[]> {
        const records = await this.model.findAll({});
        if (!records) {
            return [];
        }
        return records;
    }

    async delete(whereOptions: WhereOptions<T>): Promise<void> {
        // whereOptions is an object that contains the conditions to find the record to be deleted
        // for example, if you want to delete a record with id 1, you can
        // call this.delete({ id: 1 }) and it will delete the record with id 1.
        // if you want to delete a record with name 'John', you can call this.delete({ name: 'John' }) and it will delete the record with name 'John'.
        const record = await this.model.destroy({
            where: {
                ...whereOptions
            }
        });

        if (!record) {
            throw new NotFoundError(`Record not found for deletion with options: ${JSON.stringify(whereOptions)}`);
        }

        return;
    }

    async create(data: CreationAttributes<T>): Promise<T> {
        const record = await this.model.create(data);
        return record;
    }

    async update(id: number, data: Partial<T>): Promise<T | null> {
        const record = await this.model.findByPk(id);
        if (!record) {
            throw new NotFoundError(`Record with id ${id} not found`);
        }
        Object.assign(record, data);
        await record.save();
        return record;
    }

}

export default BaseRepository;