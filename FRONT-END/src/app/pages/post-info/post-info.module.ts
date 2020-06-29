import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { PostInfoPageRoutingModule } from './post-info-routing.module';

import { PostInfoPage } from './post-info.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    PostInfoPageRoutingModule
  ],
  declarations: [PostInfoPage]
})
export class PostInfoPageModule {}
