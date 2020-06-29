import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { TaskInfoPage } from './task-info.page';

describe('TaskInfoPage', () => {
  let component: TaskInfoPage;
  let fixture: ComponentFixture<TaskInfoPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TaskInfoPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(TaskInfoPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
